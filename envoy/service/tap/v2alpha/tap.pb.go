// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/service/tap/v2alpha/tap.proto

package envoy_service_tap_v2alpha

import (
	context "context"
	core "github.com/FeiYing9/go-control-plane/envoy/api/v2/core"
	v2alpha "github.com/FeiYing9/go-control-plane/envoy/data/tap/v2alpha"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [#not-implemented-hide:] Stream message for the Tap API. Envoy will open a stream to the server
// and stream taps without ever expecting a response.
type StreamTapsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier data effectively is a structured metadata. As a performance optimization this will
	// only be sent in the first message on the stream.
	Identifier *StreamTapsRequest_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// The trace id. this can be used to merge together a streaming trace. Note that the trace_id
	// is not guaranteed to be spatially or temporally unique.
	TraceId uint64 `protobuf:"varint,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// The trace data.
	Trace *v2alpha.TraceWrapper `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
}

func (x *StreamTapsRequest) Reset() {
	*x = StreamTapsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_tap_v2alpha_tap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamTapsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTapsRequest) ProtoMessage() {}

func (x *StreamTapsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_tap_v2alpha_tap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTapsRequest.ProtoReflect.Descriptor instead.
func (*StreamTapsRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_tap_v2alpha_tap_proto_rawDescGZIP(), []int{0}
}

func (x *StreamTapsRequest) GetIdentifier() *StreamTapsRequest_Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *StreamTapsRequest) GetTraceId() uint64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (x *StreamTapsRequest) GetTrace() *v2alpha.TraceWrapper {
	if x != nil {
		return x.Trace
	}
	return nil
}

// [#not-implemented-hide:]
type StreamTapsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamTapsResponse) Reset() {
	*x = StreamTapsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_tap_v2alpha_tap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamTapsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTapsResponse) ProtoMessage() {}

func (x *StreamTapsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_tap_v2alpha_tap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTapsResponse.ProtoReflect.Descriptor instead.
func (*StreamTapsResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_tap_v2alpha_tap_proto_rawDescGZIP(), []int{1}
}

type StreamTapsRequest_Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node sending taps over the stream.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// The opaque identifier that was set in the :ref:`output config
	// <envoy_api_field_service.tap.v2alpha.StreamingGrpcSink.tap_id>`.
	TapId string `protobuf:"bytes,2,opt,name=tap_id,json=tapId,proto3" json:"tap_id,omitempty"`
}

func (x *StreamTapsRequest_Identifier) Reset() {
	*x = StreamTapsRequest_Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_tap_v2alpha_tap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamTapsRequest_Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTapsRequest_Identifier) ProtoMessage() {}

func (x *StreamTapsRequest_Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_tap_v2alpha_tap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTapsRequest_Identifier.ProtoReflect.Descriptor instead.
func (*StreamTapsRequest_Identifier) Descriptor() ([]byte, []int) {
	return file_envoy_service_tap_v2alpha_tap_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StreamTapsRequest_Identifier) GetNode() *core.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *StreamTapsRequest_Identifier) GetTapId() string {
	if x != nil {
		return x.TapId
	}
	return ""
}

var File_envoy_service_tap_v2alpha_tap_proto protoreflect.FileDescriptor

var file_envoy_service_tap_v2alpha_tap_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x74, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74, 0x61, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a,
	0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x57, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x05, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x1a, 0x5a, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x35, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x70, 0x49, 0x64, 0x22, 0x14,
	0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x7f, 0x0a, 0x0e, 0x54, 0x61, 0x70, 0x53, 0x69, 0x6e, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x54, 0x61, 0x70, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x40, 0x0a, 0x27, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x42, 0x08, 0x54, 0x61, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x88, 0x01, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_tap_v2alpha_tap_proto_rawDescOnce sync.Once
	file_envoy_service_tap_v2alpha_tap_proto_rawDescData = file_envoy_service_tap_v2alpha_tap_proto_rawDesc
)

func file_envoy_service_tap_v2alpha_tap_proto_rawDescGZIP() []byte {
	file_envoy_service_tap_v2alpha_tap_proto_rawDescOnce.Do(func() {
		file_envoy_service_tap_v2alpha_tap_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_tap_v2alpha_tap_proto_rawDescData)
	})
	return file_envoy_service_tap_v2alpha_tap_proto_rawDescData
}

var file_envoy_service_tap_v2alpha_tap_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_service_tap_v2alpha_tap_proto_goTypes = []interface{}{
	(*StreamTapsRequest)(nil),            // 0: envoy.service.tap.v2alpha.StreamTapsRequest
	(*StreamTapsResponse)(nil),           // 1: envoy.service.tap.v2alpha.StreamTapsResponse
	(*StreamTapsRequest_Identifier)(nil), // 2: envoy.service.tap.v2alpha.StreamTapsRequest.Identifier
	(*v2alpha.TraceWrapper)(nil),         // 3: envoy.data.tap.v2alpha.TraceWrapper
	(*core.Node)(nil),                    // 4: envoy.api.v2.core.Node
}
var file_envoy_service_tap_v2alpha_tap_proto_depIdxs = []int32{
	2, // 0: envoy.service.tap.v2alpha.StreamTapsRequest.identifier:type_name -> envoy.service.tap.v2alpha.StreamTapsRequest.Identifier
	3, // 1: envoy.service.tap.v2alpha.StreamTapsRequest.trace:type_name -> envoy.data.tap.v2alpha.TraceWrapper
	4, // 2: envoy.service.tap.v2alpha.StreamTapsRequest.Identifier.node:type_name -> envoy.api.v2.core.Node
	0, // 3: envoy.service.tap.v2alpha.TapSinkService.StreamTaps:input_type -> envoy.service.tap.v2alpha.StreamTapsRequest
	1, // 4: envoy.service.tap.v2alpha.TapSinkService.StreamTaps:output_type -> envoy.service.tap.v2alpha.StreamTapsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_service_tap_v2alpha_tap_proto_init() }
func file_envoy_service_tap_v2alpha_tap_proto_init() {
	if File_envoy_service_tap_v2alpha_tap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_tap_v2alpha_tap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamTapsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_service_tap_v2alpha_tap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamTapsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_service_tap_v2alpha_tap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamTapsRequest_Identifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_service_tap_v2alpha_tap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_tap_v2alpha_tap_proto_goTypes,
		DependencyIndexes: file_envoy_service_tap_v2alpha_tap_proto_depIdxs,
		MessageInfos:      file_envoy_service_tap_v2alpha_tap_proto_msgTypes,
	}.Build()
	File_envoy_service_tap_v2alpha_tap_proto = out.File
	file_envoy_service_tap_v2alpha_tap_proto_rawDesc = nil
	file_envoy_service_tap_v2alpha_tap_proto_goTypes = nil
	file_envoy_service_tap_v2alpha_tap_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TapSinkServiceClient is the client API for TapSinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapSinkServiceClient interface {
	// Envoy will connect and send StreamTapsRequest messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect.
	StreamTaps(ctx context.Context, opts ...grpc.CallOption) (TapSinkService_StreamTapsClient, error)
}

type tapSinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTapSinkServiceClient(cc grpc.ClientConnInterface) TapSinkServiceClient {
	return &tapSinkServiceClient{cc}
}

func (c *tapSinkServiceClient) StreamTaps(ctx context.Context, opts ...grpc.CallOption) (TapSinkService_StreamTapsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapSinkService_serviceDesc.Streams[0], "/envoy.service.tap.v2alpha.TapSinkService/StreamTaps", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapSinkServiceStreamTapsClient{stream}
	return x, nil
}

type TapSinkService_StreamTapsClient interface {
	Send(*StreamTapsRequest) error
	CloseAndRecv() (*StreamTapsResponse, error)
	grpc.ClientStream
}

type tapSinkServiceStreamTapsClient struct {
	grpc.ClientStream
}

func (x *tapSinkServiceStreamTapsClient) Send(m *StreamTapsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapSinkServiceStreamTapsClient) CloseAndRecv() (*StreamTapsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTapsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TapSinkServiceServer is the server API for TapSinkService service.
type TapSinkServiceServer interface {
	// Envoy will connect and send StreamTapsRequest messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect.
	StreamTaps(TapSinkService_StreamTapsServer) error
}

// UnimplementedTapSinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapSinkServiceServer struct {
}

func (*UnimplementedTapSinkServiceServer) StreamTaps(TapSinkService_StreamTapsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTaps not implemented")
}

func RegisterTapSinkServiceServer(s *grpc.Server, srv TapSinkServiceServer) {
	s.RegisterService(&_TapSinkService_serviceDesc, srv)
}

func _TapSinkService_StreamTaps_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapSinkServiceServer).StreamTaps(&tapSinkServiceStreamTapsServer{stream})
}

type TapSinkService_StreamTapsServer interface {
	SendAndClose(*StreamTapsResponse) error
	Recv() (*StreamTapsRequest, error)
	grpc.ServerStream
}

type tapSinkServiceStreamTapsServer struct {
	grpc.ServerStream
}

func (x *tapSinkServiceStreamTapsServer) SendAndClose(m *StreamTapsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapSinkServiceStreamTapsServer) Recv() (*StreamTapsRequest, error) {
	m := new(StreamTapsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TapSinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v2alpha.TapSinkService",
	HandlerType: (*TapSinkServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTaps",
			Handler:       _TapSinkService_StreamTaps_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v2alpha/tap.proto",
}
