// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/http/original_ip_detection/custom_header/v3/custom_header.proto

package envoy_extensions_http_original_ip_detection_custom_header_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
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

// This extension allows for the original downstream remote IP to be detected
// by reading the value from a configured header name. If the value is successfully parsed
// as an IP, it'll be treated as the effective downstream remote address and seen as such
// by all filters. See :ref:`original_ip_detection_extensions
// <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.original_ip_detection_extensions>`
// for an overview of how extensions operate and what happens when an extension fails
// to detect the remote IP.
//
// [#extension: envoy.http.original_ip_detection.custom_header]
type CustomHeaderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The header name containing the original downstream remote address, if present.
	//
	// Note: in the case of a multi-valued header, only the first value is tried and the rest are ignored.
	HeaderName string `protobuf:"bytes,1,opt,name=header_name,json=headerName,proto3" json:"header_name,omitempty"`
	// If set to true, the extension could decide that the detected address should be treated as
	// trusted by the HCM. If the address is considered :ref:`trusted<config_http_conn_man_headers_x-forwarded-for_trusted_client_address>`,
	// it might be used as input to determine if the request is internal (among other things).
	AllowExtensionToSetAddressAsTrusted bool `protobuf:"varint,2,opt,name=allow_extension_to_set_address_as_trusted,json=allowExtensionToSetAddressAsTrusted,proto3" json:"allow_extension_to_set_address_as_trusted,omitempty"`
	// If this is set, the request will be rejected when detection fails using it as the HTTP response status.
	//
	// .. note::
	//   If this is set to < 400 or > 511, the default status 403 will be used instead.
	RejectWithStatus *v3.HttpStatus `protobuf:"bytes,3,opt,name=reject_with_status,json=rejectWithStatus,proto3" json:"reject_with_status,omitempty"`
}

func (x *CustomHeaderConfig) Reset() {
	*x = CustomHeaderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomHeaderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomHeaderConfig) ProtoMessage() {}

func (x *CustomHeaderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomHeaderConfig.ProtoReflect.Descriptor instead.
func (*CustomHeaderConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDescGZIP(), []int{0}
}

func (x *CustomHeaderConfig) GetHeaderName() string {
	if x != nil {
		return x.HeaderName
	}
	return ""
}

func (x *CustomHeaderConfig) GetAllowExtensionToSetAddressAsTrusted() bool {
	if x != nil {
		return x.AllowExtensionToSetAddressAsTrusted
	}
	return false
}

func (x *CustomHeaderConfig) GetRejectWithStatus() *v3.HttpStatus {
	if x != nil {
		return x.RejectWithStatus
	}
	return nil
}

var File_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto protoreflect.FileDescriptor

var file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDesc = []byte{
	0x0a, 0x50, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x12, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2e, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x01, 0xc0, 0x01,
	0x01, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x56, 0x0a, 0x29, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x23, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41,
	0x73, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x12, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x10, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x69, 0x0a, 0x4a, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x69, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42,
	0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDescOnce sync.Once
	file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDescData = file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDesc
)

func file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDescGZIP() []byte {
	file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDescData)
	})
	return file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDescData
}

var file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_goTypes = []interface{}{
	(*CustomHeaderConfig)(nil), // 0: envoy.extensions.http.original_ip_detection.custom_header.v3.CustomHeaderConfig
	(*v3.HttpStatus)(nil),      // 1: envoy.type.v3.HttpStatus
}
var file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.http.original_ip_detection.custom_header.v3.CustomHeaderConfig.reject_with_status:type_name -> envoy.type.v3.HttpStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_init()
}
func file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_init() {
	if File_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomHeaderConfig); i {
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
			RawDescriptor: file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_msgTypes,
	}.Build()
	File_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto = out.File
	file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_rawDesc = nil
	file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_goTypes = nil
	file_envoy_extensions_http_original_ip_detection_custom_header_v3_custom_header_proto_depIdxs = nil
}
