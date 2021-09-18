// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/config/filter/listener/original_src/v2alpha1/original_src.proto

package envoy_config_filter_listener_original_src_v2alpha1

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// The Original Src filter binds upstream connections to the original source address determined
// for the connection. This address could come from something like the Proxy Protocol filter, or it
// could come from trusted http headers.
type OriginalSrc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to bind the port to the one used in the original downstream connection.
	// [#not-implemented-hide:]
	BindPort bool `protobuf:"varint,1,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// Sets the SO_MARK option on the upstream connection's socket to the provided value. Used to
	// ensure that non-local addresses may be routed back through envoy when binding to the original
	// source address. The option will not be applied if the mark is 0.
	Mark uint32 `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *OriginalSrc) Reset() {
	*x = OriginalSrc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalSrc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalSrc) ProtoMessage() {}

func (x *OriginalSrc) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalSrc.ProtoReflect.Descriptor instead.
func (*OriginalSrc) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDescGZIP(), []int{0}
}

func (x *OriginalSrc) GetBindPort() bool {
	if x != nil {
		return x.BindPort
	}
	return false
}

func (x *OriginalSrc) GetMark() uint32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

var File_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto protoreflect.FileDescriptor

var file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDesc = []byte{
	0x0a, 0x45, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x72, 0x63, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x72,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x72, 0x63, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0b, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x53,
	0x72, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d,
	0x61, 0x72, 0x6b, 0x42, 0x97, 0x01, 0x0a, 0x40, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x72, 0x63, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x53, 0x72, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f,
	0x05, 0x33, 0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x72, 0x63, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDescOnce sync.Once
	file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDescData = file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDesc
)

func file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDescData)
	})
	return file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDescData
}

var file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_goTypes = []interface{}{
	(*OriginalSrc)(nil), // 0: envoy.config.filter.listener.original_src.v2alpha1.OriginalSrc
}
var file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_init() }
func file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_init() {
	if File_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginalSrc); i {
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
			RawDescriptor: file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto = out.File
	file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_rawDesc = nil
	file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_goTypes = nil
	file_envoy_config_filter_listener_original_src_v2alpha1_original_src_proto_depIdxs = nil
}
