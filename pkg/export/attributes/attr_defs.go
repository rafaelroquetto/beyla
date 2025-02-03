package attributes

import (
	"maps"

	attr "github.com/grafana/beyla/pkg/export/attributes/names"
)

// AttrGroups will let enabling by default some groups of attributes under
// given circumstances. For example, will let enabling kubernetes metadata attributes
// only if Beyla is running under Kubernetes and kube metadata is enabled.
type AttrGroups int

const (
	GroupKubernetes = AttrGroups(1 << iota)
	GroupPrometheus
	GroupHTTPRoutes
	GroupNetIfaceDirection
	GroupNetCIDR
	GroupTraces
)

func (e *AttrGroups) Has(groups AttrGroups) bool {
	return *e&groups != 0
}

func (e *AttrGroups) Add(groups AttrGroups) {
	*e |= groups
}

// Any new metric and attribute must be added here to be matched from the user-provided wildcard
// selectors of the attributes.select section
func getDefinitions(groups AttrGroups) map[Section]AttrReportGroup {
	kubeEnabled := groups.Has(GroupKubernetes)
	promEnabled := groups.Has(GroupPrometheus)
	ifaceDirEnabled := groups.Has(GroupNetIfaceDirection)
	cidrEnabled := groups.Has(GroupNetCIDR)

	// attributes to be reported exclusively for prometheus exporters
	var prometheusAttributes = AttrReportGroup{
		Disabled: !promEnabled,
		Attributes: map[attr.Name]Default{
			attr.Instance:         true,
			attr.Job:              true,
			attr.ServiceNamespace: true,
		},
	}
	// ServiceName and ServiceNamespace are reported both as resource and metric attributes, as
	// the OTEL definition requires that it is reported as resource attribute,
	// but Grafana Cloud takes it from the metric
	var appAttributes = AttrReportGroup{
		SubGroups: []*AttrReportGroup{&prometheusAttributes},
		Attributes: map[attr.Name]Default{
			attr.ServiceName:      true,
			attr.ServiceNamespace: true,
		},
	}

	// attributes to be reported exclusively for network metrics when
	// kubernetes metadata is enabled
	var networkKubeAttributes = AttrReportGroup{
		Disabled: !kubeEnabled,
		Attributes: map[attr.Name]Default{
			attr.K8sSrcOwnerName: true,
			attr.K8sSrcOwnerType: true,
			attr.K8sSrcNamespace: true,
			attr.K8sDstOwnerName: true,
			attr.K8sDstOwnerType: true,
			attr.K8sDstNamespace: true,
			attr.K8sClusterName:  true,
			attr.K8sSrcName:      false,
			attr.K8sSrcType:      false,
			attr.K8sSrcNodeIP:    false,
			attr.K8sSrcNodeName:  false,
			attr.K8sDstName:      false,
			attr.K8sDstType:      false,
			attr.K8sDstNodeIP:    false,
			attr.K8sDstNodeName:  false,
		},
	}

	// network CIDR attributes are only enabled if the CIDRs configuration
	// is defined
	var networkCIDR = AttrReportGroup{
		Disabled: !cidrEnabled,
		Attributes: map[attr.Name]Default{
			attr.DstCIDR: true,
			attr.SrcCIDR: true,
		},
	}

	// attributes to be reported exclusively for application metrics when
	// kubernetes metadata is enabled
	var appKubeAttributes = AttrReportGroup{
		Disabled: !kubeEnabled,
		Attributes: map[attr.Name]Default{
			attr.K8sNamespaceName:   true,
			attr.K8sPodName:         true,
			attr.K8sContainerName:   true,
			attr.K8sDeploymentName:  true,
			attr.K8sReplicaSetName:  true,
			attr.K8sDaemonSetName:   true,
			attr.K8sStatefulSetName: true,
			attr.K8sNodeName:        true,
			attr.K8sPodUID:          true,
			attr.K8sPodStartTime:    true,
			attr.K8sClusterName:     true,
			attr.K8sOwnerName:       true,
		},
	}

	var httpRoutes = AttrReportGroup{
		Disabled: !groups.Has(GroupHTTPRoutes),
		Attributes: map[attr.Name]Default{
			attr.HTTPRoute: true,
		},
	}

	var serverInfo = AttrReportGroup{
		Attributes: map[attr.Name]Default{
			attr.ClientAddr: false,
			attr.ServerAddr: true,
			attr.ServerPort: true,
		},
	}
	var httpClientInfo = AttrReportGroup{
		Attributes: map[attr.Name]Default{
			attr.ServerAddr: true,
			attr.ServerPort: true,
		},
	}
	var grpcClientInfo = AttrReportGroup{
		Attributes: map[attr.Name]Default{
			attr.ServerAddr: true,
		},
	}

	var httpCommon = AttrReportGroup{
		SubGroups: []*AttrReportGroup{&httpRoutes},
		Attributes: map[attr.Name]Default{
			attr.HTTPRequestMethod:      true,
			attr.HTTPResponseStatusCode: true,
			attr.HTTPUrlPath:            false,
		},
	}

	// TODO: populate it with host resource attributes in https://opentelemetry.io/docs/specs/semconv/resource/host/
	var hostAttributes = AttrReportGroup{
		Attributes: map[attr.Name]Default{
			attr.HostName: true,
		},
	}

	// the following attributes are only reported as metric attributes in Prometheus,
	// as the OTEL standard defines them as resource attributes.
	var promProcessAttributes = AttrReportGroup{
		Disabled: !promEnabled,
		Attributes: map[attr.Name]Default{
			attr.Instance:        true,
			attr.Job:             true,
			attr.ProcCommand:     true,
			attr.ProcOwner:       true,
			attr.ProcParentPid:   true,
			attr.ProcPid:         true,
			attr.ProcCommandLine: false,
			attr.ProcCommandArgs: false,
			attr.ProcExecName:    false,
			attr.ProcExecPath:    false,
		},
	}

	var processAttributes = AttrReportGroup{
		SubGroups: []*AttrReportGroup{&appKubeAttributes, &hostAttributes, &promProcessAttributes},
		Attributes: map[attr.Name]Default{
			attr.ProcCPUMode:   true,
			attr.ProcDiskIODir: true,
			attr.ProcNetIODir:  true,
		},
	}

	var messagingAttributes = AttrReportGroup{
		SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes},
		Attributes: map[attr.Name]Default{
			attr.MessagingSystem:      true,
			attr.MessagingDestination: true,
		},
	}

	return map[Section]AttrReportGroup{
		BeylaNetworkFlow.Section: {
			SubGroups: []*AttrReportGroup{&networkCIDR, &networkKubeAttributes},
			Attributes: map[attr.Name]Default{
				attr.Direction:      true,
				attr.BeylaIP:        false,
				attr.Transport:      false,
				attr.SrcAddress:     false,
				attr.DstAddres:      false,
				attr.SrcPort:        false,
				attr.DstPort:        false,
				attr.SrcName:        false,
				attr.DstName:        false,
				attr.ServerPort:     false,
				attr.ClientPort:     false,
				attr.SrcZone:        false,
				attr.DstZone:        false,
				attr.IfaceDirection: Default(ifaceDirEnabled),
				attr.Iface:          Default(ifaceDirEnabled),
			},
		},
		HTTPServerDuration.Section: {
			SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes, &httpCommon, &serverInfo},
		},
		HTTPServerRequestSize.Section: {
			SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes, &httpCommon, &serverInfo},
		},
		HTTPClientDuration.Section: {
			SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes, &httpCommon, &httpClientInfo},
		},
		HTTPClientRequestSize.Section: {
			SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes, &httpCommon, &httpClientInfo},
		},
		RPCClientDuration.Section: {
			SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes, &grpcClientInfo},
			Attributes: map[attr.Name]Default{
				attr.RPCMethod:         true,
				attr.RPCSystem:         true,
				attr.RPCGRPCStatusCode: true,
			},
		},
		RPCServerDuration.Section: {
			SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes, &serverInfo},
			Attributes: map[attr.Name]Default{
				attr.RPCMethod:         true,
				attr.RPCSystem:         true,
				attr.RPCGRPCStatusCode: true,
			},
		},
		DBClientDuration.Section: {
			SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes},
			Attributes: map[attr.Name]Default{
				attr.DBOperation:  true,
				attr.DBSystemName: true,
				attr.ErrorType:    true,
			},
		},
		MessagingPublishDuration.Section: {
			SubGroups: []*AttrReportGroup{&messagingAttributes},
		},
		MessagingProcessDuration.Section: {
			SubGroups: []*AttrReportGroup{&messagingAttributes},
		},
		Traces.Section: {
			Attributes: map[attr.Name]Default{
				attr.DBQueryText: false,
			},
		},
		ProcessCPUUtilization.Section: {SubGroups: []*AttrReportGroup{&processAttributes}},
		ProcessCPUTime.Section:        {SubGroups: []*AttrReportGroup{&processAttributes}},
		ProcessMemoryUsage.Section:    {SubGroups: []*AttrReportGroup{&processAttributes}},
		ProcessMemoryVirtual.Section:  {SubGroups: []*AttrReportGroup{&processAttributes}},
		ProcessDiskIO.Section:         {SubGroups: []*AttrReportGroup{&processAttributes}},
		ProcessNetIO.Section:          {SubGroups: []*AttrReportGroup{&processAttributes}},
		GPUKernelLaunchCalls.Section: {
			SubGroups: []*AttrReportGroup{&appAttributes, &appKubeAttributes},
			Attributes: map[attr.Name]Default{
				attr.CudaKernelName: true,
			},
		},
		GPUMemoryAllocations.Section: {
			SubGroups:  []*AttrReportGroup{&appAttributes, &appKubeAttributes},
			Attributes: map[attr.Name]Default{},
		},
		// span and service graph metrics don't yet implement attribute selection,
		// but their values can still be filtered, so we list them here just to
		// make the filter recognize its attributes
		// TODO: when service graph and spam metrics implement attribute selection, replace this section by proper metric names
		"---- temporary placeholder for span and service graph metrics ----": {
			Attributes: map[attr.Name]Default{
				attr.Client:            false,
				attr.ClientNamespace:   false,
				attr.Server:            false,
				attr.ServerNamespace:   false,
				attr.Source:            false,
				attr.Service:           false,
				attr.ServiceInstanceID: false,
				attr.ServiceNamespace:  false,
				attr.SpanKind:          false,
				attr.SpanName:          false,
				attr.StatusCode:        false,
			},
		},
	}
}

// AllAttributeNames returns a set with all the names in the attributes database
// as returned by the getDefinitions function
func AllAttributeNames() map[attr.Name]struct{} {
	names := map[attr.Name]struct{}{}
	// -1 to enable all the metric group flags
	for _, section := range getDefinitions(-1) {
		maps.Copy(names, section.All())
	}
	return names
}
