// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/config/listener/v3/quic_config.proto

package envoy_config_listener_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Configuration specific to the UDP QUIC listener.
// [#next-free-field: 8]
type QuicProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuicProtocolOptions *v3.QuicProtocolOptions `protobuf:"bytes,1,opt,name=quic_protocol_options,json=quicProtocolOptions,proto3" json:"quic_protocol_options,omitempty"`
	// Maximum number of milliseconds that connection will be alive when there is
	// no network activity. 300000ms if not specified.
	IdleTimeout *duration.Duration `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// Connection timeout in milliseconds before the crypto handshake is finished.
	// 20000ms if not specified.
	CryptoHandshakeTimeout *duration.Duration `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
	// Runtime flag that controls whether the listener is enabled or not. If not specified, defaults
	// to enabled.
	Enabled *v3.RuntimeFeatureFlag `protobuf:"bytes,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// A multiplier to number of connections which is used to determine how many packets to read per
	// event loop. A reasonable number should allow the listener to process enough payload but not
	// starve TCP and other UDP sockets and also prevent long event loop duration.
	// The default value is 32. This means if there are N QUIC connections, the total number of
	// packets to read in each read event will be 32 * N.
	// The actual number of packets to read in total by the UDP listener is also
	// bound by 6000, regardless of this field or how many connections there are.
	PacketsToReadToConnectionCountRatio *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=packets_to_read_to_connection_count_ratio,json=packetsToReadToConnectionCountRatio,proto3" json:"packets_to_read_to_connection_count_ratio,omitempty"`
	// Configure which implementation of `quic::QuicCryptoClientStreamBase` to be used for this listener.
	// If not specified the :ref:`QUICHE default one configured by <envoy_v3_api_msg_extensions.quic.crypto_stream.v3.CryptoServerStreamConfig>` will be used.
	// [#extension-category: envoy.quic.server.crypto_stream]
	CryptoStreamConfig *v3.TypedExtensionConfig `protobuf:"bytes,6,opt,name=crypto_stream_config,json=cryptoStreamConfig,proto3" json:"crypto_stream_config,omitempty"`
	// Configure which implementation of `quic::ProofSource` to be used for this listener.
	// If not specified the :ref:`default one configured by <envoy_v3_api_msg_extensions.quic.proof_source.v3.ProofSourceConfig>` will be used.
	// [#extension-category: envoy.quic.proof_source]
	ProofSourceConfig *v3.TypedExtensionConfig `protobuf:"bytes,7,opt,name=proof_source_config,json=proofSourceConfig,proto3" json:"proof_source_config,omitempty"`
}

func (x *QuicProtocolOptions) Reset() {
	*x = QuicProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v3_quic_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuicProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuicProtocolOptions) ProtoMessage() {}

func (x *QuicProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v3_quic_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuicProtocolOptions.ProtoReflect.Descriptor instead.
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v3_quic_config_proto_rawDescGZIP(), []int{0}
}

func (x *QuicProtocolOptions) GetQuicProtocolOptions() *v3.QuicProtocolOptions {
	if x != nil {
		return x.QuicProtocolOptions
	}
	return nil
}

func (x *QuicProtocolOptions) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *QuicProtocolOptions) GetCryptoHandshakeTimeout() *duration.Duration {
	if x != nil {
		return x.CryptoHandshakeTimeout
	}
	return nil
}

func (x *QuicProtocolOptions) GetEnabled() *v3.RuntimeFeatureFlag {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *QuicProtocolOptions) GetPacketsToReadToConnectionCountRatio() *wrappers.UInt32Value {
	if x != nil {
		return x.PacketsToReadToConnectionCountRatio
	}
	return nil
}

func (x *QuicProtocolOptions) GetCryptoStreamConfig() *v3.TypedExtensionConfig {
	if x != nil {
		return x.CryptoStreamConfig
	}
	return nil
}

func (x *QuicProtocolOptions) GetProofSourceConfig() *v3.TypedExtensionConfig {
	if x != nil {
		return x.ProofSourceConfig
	}
	return nil
}

var File_envoy_config_listener_v3_quic_config_proto protoreflect.FileDescriptor

var file_envoy_config_listener_v3_quic_config_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x05,
	0x0a, 0x13, 0x51, 0x75, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a, 0x15, 0x71, 0x75, 0x69, 0x63, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x75, 0x69, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x13, 0x71, 0x75, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x53, 0x0a, 0x18, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x68, 0x61, 0x6e,
	0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x16, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c,
	0x61, 0x67, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x7d, 0x0a, 0x29, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74,
	0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x23, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x54, 0x6f,
	0x52, 0x65, 0x61, 0x64, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x5c, 0x0a, 0x14, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x12, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x3a, 0x30, 0x9a, 0xc5, 0x88, 0x1e, 0x2b, 0x0a, 0x29, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x51, 0x75, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x43, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x42, 0x0f, 0x51, 0x75, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_listener_v3_quic_config_proto_rawDescOnce sync.Once
	file_envoy_config_listener_v3_quic_config_proto_rawDescData = file_envoy_config_listener_v3_quic_config_proto_rawDesc
)

func file_envoy_config_listener_v3_quic_config_proto_rawDescGZIP() []byte {
	file_envoy_config_listener_v3_quic_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_listener_v3_quic_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_listener_v3_quic_config_proto_rawDescData)
	})
	return file_envoy_config_listener_v3_quic_config_proto_rawDescData
}

var file_envoy_config_listener_v3_quic_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_listener_v3_quic_config_proto_goTypes = []interface{}{
	(*QuicProtocolOptions)(nil),     // 0: envoy.config.listener.v3.QuicProtocolOptions
	(*v3.QuicProtocolOptions)(nil),  // 1: envoy.config.core.v3.QuicProtocolOptions
	(*duration.Duration)(nil),       // 2: google.protobuf.Duration
	(*v3.RuntimeFeatureFlag)(nil),   // 3: envoy.config.core.v3.RuntimeFeatureFlag
	(*wrappers.UInt32Value)(nil),    // 4: google.protobuf.UInt32Value
	(*v3.TypedExtensionConfig)(nil), // 5: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_config_listener_v3_quic_config_proto_depIdxs = []int32{
	1, // 0: envoy.config.listener.v3.QuicProtocolOptions.quic_protocol_options:type_name -> envoy.config.core.v3.QuicProtocolOptions
	2, // 1: envoy.config.listener.v3.QuicProtocolOptions.idle_timeout:type_name -> google.protobuf.Duration
	2, // 2: envoy.config.listener.v3.QuicProtocolOptions.crypto_handshake_timeout:type_name -> google.protobuf.Duration
	3, // 3: envoy.config.listener.v3.QuicProtocolOptions.enabled:type_name -> envoy.config.core.v3.RuntimeFeatureFlag
	4, // 4: envoy.config.listener.v3.QuicProtocolOptions.packets_to_read_to_connection_count_ratio:type_name -> google.protobuf.UInt32Value
	5, // 5: envoy.config.listener.v3.QuicProtocolOptions.crypto_stream_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	5, // 6: envoy.config.listener.v3.QuicProtocolOptions.proof_source_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_config_listener_v3_quic_config_proto_init() }
func file_envoy_config_listener_v3_quic_config_proto_init() {
	if File_envoy_config_listener_v3_quic_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_listener_v3_quic_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuicProtocolOptions); i {
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
			RawDescriptor: file_envoy_config_listener_v3_quic_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_listener_v3_quic_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_listener_v3_quic_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_listener_v3_quic_config_proto_msgTypes,
	}.Build()
	File_envoy_config_listener_v3_quic_config_proto = out.File
	file_envoy_config_listener_v3_quic_config_proto_rawDesc = nil
	file_envoy_config_listener_v3_quic_config_proto_goTypes = nil
	file_envoy_config_listener_v3_quic_config_proto_depIdxs = nil
}
