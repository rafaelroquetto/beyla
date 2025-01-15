package ebpfcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/http2"
)

func TestHTTP2QuickDetection(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		inputLen int
		expected bool
	}{
		{
			name:     "Status instead of start",
			input:    []byte{0, 0, 29, 1, 4, 0, 0, 1, 101, 136, 224, 223, 222, 221, 97, 150, 223, 105, 126, 148, 19, 106, 101, 182, 165, 4, 1, 52, 160, 92, 184, 23, 174, 1, 197, 49, 104, 223, 0, 0, 44, 0, 0, 0, 0, 1, 101, 1, 0, 0, 0, 39, 31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 18, 98, 11, 14, 113, 12, 241, 116, 150, 98, 206, 79, 75, 83, 98, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 100,
			expected: true,
		},
		{
			name:     "Empty",
			input:    []byte{},
			inputLen: 100,
			expected: false,
		},
		{
			name:     "Short",
			input:    []byte{0, 0, 70, 1, 4},
			inputLen: 3,
			expected: false,
		},
		{
			name:     "Regular HTTP2/gRPC Frame",
			input:    []byte{0, 0, 70, 1, 4, 0, 0, 0, 19, 204, 131, 4, 147, 96, 233, 45, 18, 22, 147, 175, 12, 155, 139, 103, 115, 16, 172, 98, 42, 97, 145, 31, 134, 126, 167, 0, 22, 16, 7, 36, 140, 179, 27, 50, 202, 25, 101, 105, 182, 93, 33, 66, 211, 97, 41, 64, 0, 182, 66, 44, 219, 242, 186, 217, 2, 203, 196, 3, 143, 182, 209, 86, 0, 127, 203, 202, 201, 200, 199, 0, 0, 5, 0, 0, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 10000,
			expected: true,
		},
		{
			name:     "Reset frame before HTTP2/gRPC Frame",
			input:    []byte{0, 0, 4, 3, 0, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 70, 1, 4, 0, 0, 0, 21, 205, 131, 4, 147, 96, 233, 45, 18, 22, 147, 175, 12, 155, 139, 103, 115, 16, 172, 98, 42, 97, 145, 31, 134, 126, 167, 0, 22, 44, 99, 27, 33, 124, 174, 72, 228, 109, 129, 233, 27, 125, 246, 133, 44, 101, 28, 111, 70, 32, 178, 85, 163, 108, 97, 149, 199, 99, 121, 169, 90, 149, 225, 188, 176, 3, 204, 203, 202, 201, 200, 0, 0, 5, 0, 0, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 10000,
			expected: true,
		},
		{
			name:     "Too short of input len, but enough to parse the reset frame",
			input:    []byte{0, 0, 4, 3, 0, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 70, 1, 4, 0, 0, 0, 21, 205, 131, 4, 147, 96, 233, 45, 18, 22, 147, 175, 12, 155, 139, 103, 115, 16, 172, 98, 42, 97, 145, 31, 134, 126, 167, 0, 22, 44, 99, 27, 33, 124, 174, 72, 228, 109, 129, 233, 27, 125, 246, 133, 44, 101, 28, 111, 70, 32, 178, 85, 163, 108, 97, 149, 199, 99, 121, 169, 90, 149, 225, 188, 176, 3, 204, 203, 202, 201, 200, 0, 0, 5, 0, 0, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: frameHeaderLen + 2,
			expected: false,
		},
		{
			name:     "Kafka frame instead of HTTP2",
			input:    []byte{0, 0, 0, 1, 0, 0, 0, 7, 0, 0, 0, 2, 0, 6, 115, 97, 114, 97, 109, 97, 255, 255, 255, 255, 0, 0, 39, 16, 0, 0, 0, 1, 0, 9, 105, 109, 112, 111, 114, 116, 97, 110, 116, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 72},
			inputLen: 10000,
			expected: false,
		},
		{
			name:     "No headers frame (manually tweaked the type to fail)",
			input:    []byte{0, 0, 4, 3, 0, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 70, 2, 4, 0, 0, 0, 21, 205, 131, 4, 147, 96, 233, 45, 18, 22, 147, 175, 12, 155, 139, 103, 115, 16, 172, 98, 42, 97, 145, 31, 134, 126, 167, 0, 22, 44, 99, 27, 33, 124, 174, 72, 228, 109, 129, 233, 27, 125, 246, 133, 44, 101, 28, 111, 70, 32, 178, 85, 163, 108, 97, 149, 199, 99, 121, 169, 90, 149, 225, 188, 176, 3, 204, 203, 202, 201, 200, 0, 0, 5, 0, 0, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 10000,
			expected: false,
		},
		{
			name:     "Truncated frame, len should be 70 of the second frame",
			input:    []byte{0, 0, 4, 3, 0, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 70, 2, 4, 0, 0, 0, 21, 205, 131},
			inputLen: 10000,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isLikelyHTTP2(tt.input, tt.inputLen)
			assert.Equal(t, tt.expected, res)
			res1 := isHTTP2(tt.input, tt.inputLen)
			assert.Equal(t, tt.expected, res1)
		})
	}
}

func TestHTTP2Parsing(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		inputLen    int
		method      string
		path        string
		contentType string
	}{
		{
			name:        "One",
			input:       []byte{0, 0, 88, 1, 4, 0, 0, 6, 237, 208, 131, 4, 164, 96, 233, 45, 18, 22, 147, 175, 180, 164, 61, 52, 150, 169, 6, 147, 30, 173, 197, 179, 37, 2, 0, 0, 0, 0, 0, 0, 187, 70, 76, 66, 163, 126, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 213, 255, 255, 255, 255, 255, 255, 255, 1, 105, 108, 100, 108, 105, 102, 101, 0, 0, 0, 0, 0, 0, 0, 0, 64, 183, 2, 212, 164, 126, 0, 0, 64, 183, 2, 212, 164, 126, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 60, 103, 110, 32, 119, 105, 108, 108, 32, 119, 105, 116, 104, 115, 116, 97, 110, 100, 32, 96, 32, 0, 196, 164, 126, 0, 0, 60, 0, 0, 0, 0, 0, 0, 0, 112, 32, 0, 196, 164, 126, 0, 0, 137, 42, 109, 81, 165, 126, 0, 0, 97, 115, 104, 108, 105, 103, 104, 116, 46, 106, 112, 103, 42, 12, 10, 3, 85, 83, 68, 16, 57, 24, 128, 232, 146, 38, 50, 11, 97, 99, 99, 101, 115, 115, 111, 114, 105, 101, 115, 50, 11, 102, 108, 97, 115, 104, 108, 105, 103, 104, 116, 115, 10, 165, 5, 10},
			inputLen:    32,
			method:      "POST",
			path:        "",
			contentType: "",
		},
		{
			name:        "Two",
			input:       []byte{0, 0, 77, 1, 4, 0, 0, 0, 37, 195, 194, 131, 134, 193, 192, 191, 190, 0, 11, 116, 114, 97, 99, 0, 0, 0, 0, 0, 0, 0, 0, 8, 101, 112, 97, 114, 101, 110, 116, 55, 0, 8, 6, 0, 0, 0, 0, 0, 36, 42, 35, 123, 242, 89, 199, 0, 7, 1, 240, 184, 117, 0, 0, 55, 0, 0, 0, 0, 0, 0, 0, 16, 7, 1, 240, 184, 117, 0, 0, 137, 218, 220, 116, 185, 117, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 23, 0, 0, 4, 8, 0, 0, 0, 0, 37, 0, 0, 0, 5, 0, 0, 5, 0, 1, 0, 0, 0, 37, 0, 0, 0, 0, 0, 0, 0, 0, 0, 17, 0, 0, 0, 0, 0, 0, 4, 8, 0, 0, 0, 0, 0, 0, 0, 20, 12, 0, 240, 184, 117, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 210, 202, 123, 115, 185, 117, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 31, 0, 240, 184, 117, 0, 0, 174, 233, 21, 115, 185, 117, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 96, 65, 0, 240, 184, 117, 0, 0, 31, 0, 0, 0, 0, 0, 0, 0, 112, 65, 0, 240, 184, 117, 0, 0, 208, 201, 127, 3, 185, 117, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			inputLen:    126,
			method:      "POST",
			path:        "",
			contentType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			framer := byteFramer(tt.input[:tt.inputLen])
			for {
				f, err := framer.ReadFrame()

				if err != nil {
					break
				}

				if ff, ok := f.(*http2.HeadersFrame); ok {
					connInfo := BPFConnInfo{}
					method, path, contentType, _ := readMetaFrame(&connInfo, false, framer, ff)
					assert.Equal(t, tt.method, method)
					assert.Equal(t, tt.path, path)
					assert.Equal(t, tt.contentType, contentType)
				}
			}
		})
	}
}

func TestHTTP2EventsParsing(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		rinput   []byte
		inputLen int
		ignored  bool
	}{
		{
			name:     "Ignored, buffers reversed, nothing in there",
			input:    []byte{0, 0, 6, 1, 4, 0, 0, 0, 11, 136, 196, 195, 194, 193, 190, 150, 223, 105, 126, 148, 19, 106, 101, 182, 165, 4, 1, 52, 160, 94, 184, 39, 46, 52, 242, 152, 180, 111, 255, 18, 98, 11, 14, 113, 12, 241, 116, 150, 98, 206, 79, 75, 83, 98, 0, 4, 0, 0, 255, 255, 211, 196, 47, 145},
			rinput:   []byte{0, 0, 138, 1, 36, 0, 0, 0, 11, 0, 0, 0, 0, 15, 0, 0, 0, 0, 45, 0, 0, 0, 0, 0, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 201,
			ignored:  true,
		},
		{
			name:     "Not reversed",
			rinput:   []byte{0, 0, 6, 1, 4, 0, 0, 0, 11, 136, 196, 195, 194, 193, 190, 150, 223, 105, 126, 148, 19, 106, 101, 182, 165, 4, 1, 52, 160, 94, 184, 39, 46, 52, 242, 152, 180, 111, 255, 18, 98, 11, 14, 113, 12, 241, 116, 150, 98, 206, 79, 75, 83, 98, 0, 4, 0, 0, 255, 255, 211, 196, 47, 145},
			input:    []byte{0, 0, 138, 1, 36, 0, 0, 0, 11, 0, 0, 0, 0, 15, 0, 0, 0, 0, 45, 0, 0, 0, 0, 0, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 201,
			ignored:  false,
		},
		{
			name:     "New with concat",
			input:    []byte{0, 0, 138, 1, 36, 0, 0, 0, 21, 0, 0, 0, 0, 15, 222, 221, 131, 134, 220, 219, 218, 127, 0, 55, 48, 48, 45, 102, 50, 100, 52, 101, 54, 99, 98, 54, 56, 98, 53, 55, 51, 54, 56, 49, 49, 48, 99, 48, 52, 102, 49, 48, 100, 51, 101, 53, 54, 53, 56, 45, 56, 57, 57, 57, 51, 97, 48, 57, 50, 54, 51, 99, 100, 98, 49, 48, 45, 48, 49, 126, 55, 48, 48, 45, 102, 50, 100, 52, 101, 54, 99, 98, 54, 56, 98, 53, 55, 51, 54, 56, 49, 49, 48, 99, 48, 52, 102, 49, 48, 100, 51, 101, 53, 54, 53, 56, 45, 102, 49, 52, 49, 99, 49, 98, 51, 102, 57, 55, 53, 97, 49, 48, 53, 45, 48, 49, 217, 127, 1, 7, 52, 57, 57, 54, 49, 51, 117, 0, 0, 45, 0, 1, 0, 0, 0, 21, 0, 0, 0, 0, 40, 10, 16, 97, 100, 83, 101, 114, 118, 105, 99, 101, 72, 105, 103, 104, 67, 112, 117, 18, 20, 10, 18, 10, 12, 116, 97, 114, 103, 101, 116, 105, 110, 103, 75, 101, 121, 18, 2, 26, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			rinput:   []byte{0, 0, 29, 1, 4, 0, 0, 0, 21, 136, 197, 196, 195, 194, 97, 150, 228, 89, 62, 148, 19, 138, 101, 182, 165, 4, 1, 52, 160, 65, 113, 176, 220, 105, 213, 49, 104, 223, 255, 18, 226, 15, 113, 12, 114, 119, 13, 241, 244, 115, 143, 247, 117, 12, 113, 246, 144, 98, 206, 79, 75, 83, 98, 0},
			inputLen: 201,
			ignored:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := makeBPFHTTP2Info(tt.input, tt.rinput, tt.inputLen)
			_, ignore, _ := http2FromBuffers(&info)
			assert.Equal(t, tt.ignored, ignore)
		})
	}
}

func TestDynamicTableUpdates(t *testing.T) {
	rinput := []byte{0, 0, 138, 1, 36, 0, 0, 0, 11, 0, 0, 0, 0, 15, 0, 0, 0, 0, 45, 0, 0, 0, 0, 0, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	tests := []struct {
		name     string
		input    []byte
		inputLen int
	}{
		{
			name:     "Full path, lots of headers",
			input:    []byte{0, 0, 222, 1, 4, 0, 0, 0, 1, 64, 5, 58, 112, 97, 116, 104, 33, 47, 114, 111, 117, 116, 101, 103, 117, 105, 100, 101, 46, 82, 111, 117, 116, 101, 71, 117, 105, 100, 101, 47, 71, 101, 116, 70, 101, 97, 116, 117, 114, 101, 64, 10, 58, 97, 117, 116, 104, 111, 114, 105, 116, 121, 15, 108, 111, 99, 97, 108, 104, 111, 115, 116, 58, 53, 48, 48, 53, 49, 131, 134, 64, 12, 99, 111, 110, 116, 101, 110, 116, 45, 116, 121, 112, 101, 16, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 103, 114, 112, 99, 64, 2, 116, 101, 8, 116, 114, 97, 105, 108, 101, 114, 115, 64, 20, 103, 114, 112, 99, 45, 97, 99, 99, 101, 112, 116, 45, 101, 110, 99, 111, 100, 105, 110, 103, 23, 105, 100, 101, 110, 116, 105, 116, 121, 44, 32, 100, 101, 102, 108, 97, 116, 101, 44, 32, 103, 122, 105, 112, 64, 10, 117, 115, 101, 114, 45, 97, 103, 101, 110, 116, 48, 103, 114, 112, 99, 45, 112, 121, 116, 104, 111, 110, 47, 49, 46, 54, 57, 46, 48, 32, 103, 114, 112, 99, 45, 99, 47, 52, 52, 46, 50, 46, 48, 32, 40, 108, 105, 110, 117, 120, 59, 32, 99, 104, 116, 116, 112, 50, 41, 0, 0, 4, 8, 0, 0, 0, 0, 1, 0, 0, 0, 5, 0, 0, 22, 0, 1, 0, 0, 0, 1, 0, 0, 0},
			inputLen: 1024,
		},
		{
			name:     "Full path only",
			input:    []byte{0, 0, 222, 1, 4, 0, 0, 0, 1, 64, 5, 58, 112, 97, 116, 104, 33, 47, 114, 111, 117, 116, 101, 103, 117, 105, 100, 101, 46, 82, 111, 117, 116, 101, 71, 117, 105, 100, 101, 47, 71, 101, 116, 70, 101, 97, 116, 117, 114, 101, 131},
			inputLen: 1024,
		},
		{
			name:     "Index encoded",
			input:    []byte{0, 0, 8, 1, 4, 0, 0, 0, 3, 195, 194, 131, 134, 193, 192, 191, 190, 0, 0, 4, 8, 0, 0, 0, 0, 3, 0, 0, 0, 5, 0, 0, 5, 0, 1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 4, 8, 0, 0, 0, 0, 0, 0, 0, 0, 84},
			inputLen: 1024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := makeBPFHTTP2InfoNewRequest(tt.input, rinput, tt.inputLen)
			s, ignore, _ := http2FromBuffers(&info)
			assert.False(t, ignore)
			assert.Equal(t, "POST", s.Method)
			assert.Equal(t, "/routeguide.RouteGuide/GetFeature", s.Path)
		})
	}

	// Now let's break the decoder with pushing unknown indices
	unknownIndexInput := []byte{0, 0, 8, 1, 4, 0, 0, 0, 3, 199, 200, 131, 134, 201, 202, 203, 204, 0, 0, 4, 8, 0, 0, 0, 0, 3, 0, 0, 0, 5, 0, 0, 5, 0, 1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 4, 8, 0, 0, 0, 0, 0, 0, 0, 0, 84}

	info := makeBPFHTTP2InfoNewRequest(unknownIndexInput, rinput, 1024)
	s, ignore, _ := http2FromBuffers(&info)
	assert.False(t, ignore)
	assert.Equal(t, "POST", s.Method)
	assert.Equal(t, "*", s.Path)

	nextIndex := 8 + 61 // 61 is the static table index size, 7 is how many entries we store in the dynamic table with that first request

	// Now let's send new path
	newPathInput := []byte{0, 0, 222, 1, 4, 0, 0, 0, 1, 64, 5, 58, 112, 97, 116, 104, 33, 47, 112, 111, 117, 116, 101, 103, 117, 105, 100, 101, 46, 82, 111, 117, 116, 101, 71, 117, 105, 100, 101, 47, 71, 101, 116, 70, 101, 97, 116, 117, 114, 101, 64, 10, 58, 97, 117, 116, 104, 111, 114, 105, 116, 121, 15, 108, 111, 99, 97, 108, 104, 111, 115, 116, 58, 53, 48, 48, 53, 49, 131, 134, 64, 12, 99, 111, 110, 116, 101, 110, 116, 45, 116, 121, 112, 101, 16, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 103, 114, 112, 99, 64, 2, 116, 101, 8, 116, 114, 97, 105, 108, 101, 114, 115, 64, 20, 103, 114, 112, 99, 45, 97, 99, 99, 101, 112, 116, 45, 101, 110, 99, 111, 100, 105, 110, 103, 23, 105, 100, 101, 110, 116, 105, 116, 121, 44, 32, 100, 101, 102, 108, 97, 116, 101, 44, 32, 103, 122, 105, 112, 64, 10, 117, 115, 101, 114, 45, 97, 103, 101, 110, 116, 48, 103, 114, 112, 99, 45, 112, 121, 116, 104, 111, 110, 47, 49, 46, 54, 57, 46, 48, 32, 103, 114, 112, 99, 45, 99, 47, 52, 52, 46, 50, 46, 48, 32, 40, 108, 105, 110, 117, 120, 59, 32, 99, 104, 116, 116, 112, 50, 41, 0, 0, 4, 8, 0, 0, 0, 0, 1, 0, 0, 0, 5, 0, 0, 22, 0, 1, 0, 0, 0, 1, 0, 0, 0}

	// We'll be able to decode this correctly, even with broken decoder, beause the values are sent as text
	info = makeBPFHTTP2InfoNewRequest(newPathInput, rinput, 1024)
	s, ignore, _ = http2FromBuffers(&info)
	assert.False(t, ignore)
	assert.Equal(t, "POST", s.Method)
	assert.Equal(t, "/pouteguide.RouteGuide/GetFeature", s.Path) // this value is the same I just changed the first character from r to p

	// indexed version of newPathInput
	// if we cached a new pair nextIndex + 128 is the high bit encoded next index which should be in the dynamic table
	// however we mark the decoder as invalid and it shouldn't resolve to anything for :path
	indexedNewPath := []byte{0, 0, 8, 1, 4, 0, 0, 0, 3, 195, 194, 131, 134, 193, 192, 191, byte(nextIndex + 128), 0, 0, 4, 8, 0, 0, 0, 0, 3, 0, 0, 0, 5, 0, 0, 5, 0, 1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 4, 8, 0, 0, 0, 0, 0, 0, 0, 0, 84}

	info = makeBPFHTTP2InfoNewRequest(indexedNewPath, rinput, 1024)
	s, ignore, _ = http2FromBuffers(&info)
	assert.False(t, ignore)
	assert.Equal(t, "POST", s.Method)
	assert.Equal(t, "*", s.Path) // this value is the same I just changed the first character from r to p
}

func makeBPFHTTP2Info(buf, rbuf []byte, len int) BPFHTTP2Info {
	var info BPFHTTP2Info
	copy(info.Data[:], buf)
	copy(info.RetData[:], rbuf)
	info.Len = int32(len)

	return info
}

func makeBPFHTTP2InfoNewRequest(buf, rbuf []byte, len int) BPFHTTP2Info {
	info := makeBPFHTTP2Info(buf, rbuf, len)
	info.ConnInfo.D_port = 1
	info.ConnInfo.S_port = 1
	info.NewConn = 1

	return info
}
