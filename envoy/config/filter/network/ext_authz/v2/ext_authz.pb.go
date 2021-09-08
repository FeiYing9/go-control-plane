// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/config/filter/network/ext_authz/v2/ext_authz.proto

package envoy_config_filter_network_ext_authz_v2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// External Authorization filter calls out to an external service over the
// gRPC Authorization API defined by
// :ref:`CheckRequest <envoy_api_msg_service.auth.v2.CheckRequest>`.
// A failed check will cause this filter to close the TCP connection.
type ExtAuthz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The external authorization gRPC service configuration.
	// The default timeout is set to 200ms by this filter.
	GrpcService *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. When it is set to true, Envoy will also allow traffic in case of
	// communication failure between authorization service and the proxy.
	// Defaults to false.
	FailureModeAllow bool `protobuf:"varint,3,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	// Specifies if the peer certificate is sent to the external service.
	//
	// When this field is true, Envoy will include the peer X.509 certificate, if available, in the
	// :ref:`certificate<envoy_api_field_service.auth.v2.AttributeContext.Peer.certificate>`.
	IncludePeerCertificate bool `protobuf:"varint,4,opt,name=include_peer_certificate,json=includePeerCertificate,proto3" json:"include_peer_certificate,omitempty"`
}

func (x *ExtAuthz) Reset() {
	*x = ExtAuthz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtAuthz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtAuthz) ProtoMessage() {}

func (x *ExtAuthz) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtAuthz.ProtoReflect.Descriptor instead.
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDescGZIP(), []int{0}
}

func (x *ExtAuthz) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *ExtAuthz) GetGrpcService() *core.GrpcService {
	if x != nil {
		return x.GrpcService
	}
	return nil
}

func (x *ExtAuthz) GetFailureModeAllow() bool {
	if x != nil {
		return x.FailureModeAllow
	}
	return false
}

func (x *ExtAuthz) GetIncludePeerCertificate() bool {
	if x != nil {
		return x.IncludePeerCertificate
	}
	return false
}

var File_envoy_config_filter_network_ext_authz_v2_ext_authz_proto protoreflect.FileDescriptor

var file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x65, 0x78,
	0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x74, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x32, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x12,
	0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x41, 0x0a, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x65, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x42, 0x86, 0x01, 0x0a, 0x36, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x32, 0x42,
	0x0d, 0x45, 0x78, 0x74, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x2f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDescOnce sync.Once
	file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDescData = file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDesc
)

func file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDescData)
	})
	return file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDescData
}

var file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_goTypes = []interface{}{
	(*ExtAuthz)(nil),         // 0: envoy.config.filter.network.ext_authz.v2.ExtAuthz
	(*core.GrpcService)(nil), // 1: envoy.api.v2.core.GrpcService
}
var file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_depIdxs = []int32{
	1, // 0: envoy.config.filter.network.ext_authz.v2.ExtAuthz.grpc_service:type_name -> envoy.api.v2.core.GrpcService
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_init() }
func file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_init() {
	if File_envoy_config_filter_network_ext_authz_v2_ext_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtAuthz); i {
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
			RawDescriptor: file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_network_ext_authz_v2_ext_authz_proto = out.File
	file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_rawDesc = nil
	file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_goTypes = nil
	file_envoy_config_filter_network_ext_authz_v2_ext_authz_proto_depIdxs = nil
}
