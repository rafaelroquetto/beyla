#ifndef TC_SOCK_H
#define TC_SOCK_H

#include "vmlinux.h"
#include "bpf_helpers.h"
#include "bpf_endian.h"
#include "http_types.h"
#include "tc_common.h"

#define SOCKOPS_MAP_SIZE 65535

struct {
    __uint(type, BPF_MAP_TYPE_SOCKHASH);
    __uint(max_entries, SOCKOPS_MAP_SIZE);
    __uint(key_size, sizeof(connection_info_t));
    __uint(value_size, sizeof(uint32_t));
} sock_dir SEC(".maps");

typedef struct tc_http_ctx {
    u32 offset;
    u32 seen;
    u32 written;
} __attribute__((packed)) tc_http_ctx_t;

struct tc_http_ctx_map {
    __uint(type, BPF_MAP_TYPE_LRU_HASH);
    __type(key, u32);
    __type(value, struct tc_http_ctx);
    __uint(max_entries, 10240);
} tc_http_ctx_map SEC(".maps");

typedef struct msg_data {
    u8 buf[1024];
} msg_data_t;

struct {
    __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
    __type(key, int);
    __type(value, msg_data_t);
    __uint(max_entries, 1);
} buf_mem SEC(".maps");

static __always_inline void sk_ops_extract_key_ip4(struct bpf_sock_ops *ops,
                                                   connection_info_t *conn) {
    __builtin_memcpy(conn->s_addr, ip4ip6_prefix, sizeof(ip4ip6_prefix));
    conn->s_ip[3] = ops->local_ip4;
    __builtin_memcpy(conn->d_addr, ip4ip6_prefix, sizeof(ip4ip6_prefix));
    conn->d_ip[3] = ops->remote_ip4;

    conn->s_port = ops->local_port;
    conn->d_port = bpf_ntohl(ops->remote_port);
}

// I couldn't break this up into functions, ended up running into a verifier error about ctx already written
__attribute__((unused)) static __always_inline void
sk_ops_extract_key_ip6(struct bpf_sock_ops *ops, connection_info_t *conn) {
    conn->s_ip[0] = ops->local_ip6[0];
    conn->s_ip[1] = ops->local_ip6[1];
    conn->s_ip[2] = ops->local_ip6[2];
    conn->s_ip[3] = ops->local_ip6[3];
    conn->d_ip[0] = ops->remote_ip6[0];
    conn->d_ip[1] = ops->remote_ip6[1];
    conn->d_ip[2] = ops->remote_ip6[2];
    conn->d_ip[3] = ops->remote_ip6[3];

    conn->s_port = ops->local_port;
    conn->d_port = bpf_ntohl(ops->remote_port);
}

static __always_inline void sk_msg_extract_key_ip4(struct sk_msg_md *msg, connection_info_t *conn) {
    __builtin_memcpy(conn->s_addr, ip4ip6_prefix, sizeof(ip4ip6_prefix));
    conn->s_ip[3] = msg->local_ip4;
    __builtin_memcpy(conn->d_addr, ip4ip6_prefix, sizeof(ip4ip6_prefix));
    conn->d_ip[3] = msg->remote_ip4;

    conn->s_port = msg->local_port;
    conn->d_port = bpf_ntohl(msg->remote_port);
}

__attribute__((unused)) static __always_inline void
sk_msg_extract_key_ip6(struct sk_msg_md *msg, connection_info_t *conn) {
    conn->s_ip[0] = msg->local_ip6[0];
    conn->s_ip[1] = msg->local_ip6[1];
    conn->s_ip[2] = msg->local_ip6[2];
    conn->s_ip[3] = msg->local_ip6[3];
    conn->d_ip[0] = msg->remote_ip6[0];
    conn->d_ip[1] = msg->remote_ip6[1];
    conn->d_ip[2] = msg->remote_ip6[2];
    conn->d_ip[3] = msg->remote_ip6[3];

    conn->s_port = msg->local_port;
    conn->d_port = bpf_ntohl(msg->remote_port);
}

static __always_inline void bpf_sock_ops_establish_cb(struct bpf_sock_ops *skops) {
    connection_info_t conn = {};

    // if (skops->family == AF_INET6) {
    //     sk_ops_extract_key_ip6(skops, &conn);
    // } else {
    sk_ops_extract_key_ip4(skops, &conn);
    // }

    bpf_printk("SET %llx:%d -> %llx:%d", conn.s_ip[3], conn.s_port, conn.d_ip[3], conn.d_port);
    bpf_sock_hash_update(skops, &sock_dir, &conn, BPF_ANY);
}

SEC("sockops")
int sockmap_tracker(struct bpf_sock_ops *skops) {
    u32 op = skops->op;

    switch (op) {
    case BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB:
    case BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB:
        bpf_sock_ops_establish_cb(skops);
        break;
    default:
        break;
    }
    return 0;
}

static __always_inline msg_data_t *buffer() {
    int zero = 0;
    return (msg_data_t *)bpf_map_lookup_elem(&buf_mem, &zero);
}

static __always_inline u8 is_tracked(connection_info_t *conn) {
    egress_key_t e_key = {
        .d_port = conn->d_port,
        .s_port = conn->s_port,
    };

    sort_egress_key(&e_key);

    tp_info_pid_t *tp = bpf_map_lookup_elem(&outgoing_trace_map, &e_key);
    return tp != 0;
}

SEC("sk_msg")
int packet_extender(struct sk_msg_md *msg) {
    u64 len = (u64)msg->data_end - (u64)msg->data;
    connection_info_t conn = {};

    if (msg->family == AF_INET6) {
        return 0;
    }
    // if (msg->family == AF_INET6) {
    //     sk_msg_extract_key_ip6(msg, &conn);
    // } else {
    sk_msg_extract_key_ip4(msg, &conn);
    // }

    bpf_printk("MSG %llx:%d -> %llx:%d", conn.s_ip[3], conn.s_port, conn.d_ip[3], conn.d_port);

    u8 tracked = is_tracked(&conn);

    if (tracked) {
        bpf_printk("*tracked*");
    }

    if (tracked && len > 32) {
        msg_data_t *msg_data = buffer();
        if (msg_data) {
            bpf_msg_pull_data(msg, 0, 1024, 0);
            bpf_probe_read_kernel(msg_data->buf, 1024, msg->data);
            if (is_http_request_buf(msg_data->buf)) {
                bpf_printk("len %d, s_port %d, buf: %s", len, msg->local_port, msg_data->buf);

                int newline_pos = find_first_pos_of(msg_data->buf, &msg_data->buf[1023], '\n');

                if (newline_pos >= 0) {
                    newline_pos++;
                    if (!bpf_msg_push_data(msg, newline_pos, EXTEND_SIZE, 0)) {
                        tc_http_ctx_t ctx = {
                            .offset = newline_pos,
                            .seen = 0,
                            .written = 0,
                        };
                        u32 port = msg->local_port;

                        bpf_map_update_elem(&tc_http_ctx_map, &port, &ctx, BPF_ANY);
                    }
                    bpf_msg_pull_data(msg, 0, 1024, 0);
                    bpf_printk("offset %d, new data: %s", newline_pos, (char *)msg->data);
                }
            }
        }
    }

    return SK_PASS;
}

#endif