// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/extensions/filters/network/jres_proxy/v4alpha/route.proto

package envoy_extensions_filters_network_jres_proxy_v4alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/route/v4alpha"
	v4alpha1 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v4alpha"
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

// [#next-free-field: 6]
type RouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The interface name of the service.
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	// Which group does the interface belong to.
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	// The version number of the interface.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes []*Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *RouteConfiguration) Reset() {
	*x = RouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteConfiguration) ProtoMessage() {}

func (x *RouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteConfiguration.ProtoReflect.Descriptor instead.
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteConfiguration) GetInterface() string {
	if x != nil {
		return x.Interface
	}
	return ""
}

func (x *RouteConfiguration) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *RouteConfiguration) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RouteConfiguration) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetMatch() *RouteMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *Route) GetRoute() *RouteAction {
	if x != nil {
		return x.Route
	}
	return nil
}

type RouteMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Method level routing matching.
	Method *MethodMatch `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config).
	Headers []*v4alpha.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *RouteMatch) Reset() {
	*x = RouteMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatch) ProtoMessage() {}

func (x *RouteMatch) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatch.ProtoReflect.Descriptor instead.
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescGZIP(), []int{2}
}

func (x *RouteMatch) GetMethod() *MethodMatch {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *RouteMatch) GetHeaders() []*v4alpha.HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

type RouteAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
}

func (x *RouteAction) Reset() {
	*x = RouteAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteAction) ProtoMessage() {}

func (x *RouteAction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteAction.ProtoReflect.Descriptor instead.
func (*RouteAction) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescGZIP(), []int{3}
}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (x *RouteAction) GetCluster() string {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (x *RouteAction) GetWeightedClusters() *v4alpha.WeightedCluster {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	// Indicates the upstream cluster to which the request should be routed.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	// Multiple upstream clusters can be specified for a given route. The
	// request is routed to one of the upstream clusters based on weights
	// assigned to each cluster.
	// Currently ClusterWeight only supports the name and weight fields.
	WeightedClusters *v4alpha.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

type MethodMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the method.
	Name *v4alpha1.StringMatcher `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Method parameter definition.
	// The key is the parameter index, starting from 0.
	// The value is the parameter matching type.
	ParamsMatch map[uint32]*MethodMatch_ParameterMatchSpecifier `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch,proto3" json:"params_match,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MethodMatch) Reset() {
	*x = MethodMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodMatch) ProtoMessage() {}

func (x *MethodMatch) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodMatch.ProtoReflect.Descriptor instead.
func (*MethodMatch) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescGZIP(), []int{4}
}

func (x *MethodMatch) GetName() *v4alpha1.StringMatcher {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *MethodMatch) GetParamsMatch() map[uint32]*MethodMatch_ParameterMatchSpecifier {
	if x != nil {
		return x.ParamsMatch
	}
	return nil
}

// The parameter matching type.
type MethodMatch_ParameterMatchSpecifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ParameterMatchSpecifier:
	//	*MethodMatch_ParameterMatchSpecifier_ExactMatch
	//	*MethodMatch_ParameterMatchSpecifier_RangeMatch
	ParameterMatchSpecifier isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier `protobuf_oneof:"parameter_match_specifier"`
}

func (x *MethodMatch_ParameterMatchSpecifier) Reset() {
	*x = MethodMatch_ParameterMatchSpecifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodMatch_ParameterMatchSpecifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodMatch_ParameterMatchSpecifier) ProtoMessage() {}

func (x *MethodMatch_ParameterMatchSpecifier) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodMatch_ParameterMatchSpecifier.ProtoReflect.Descriptor instead.
func (*MethodMatch_ParameterMatchSpecifier) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescGZIP(), []int{4, 0}
}

func (m *MethodMatch_ParameterMatchSpecifier) GetParameterMatchSpecifier() isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier {
	if m != nil {
		return m.ParameterMatchSpecifier
	}
	return nil
}

func (x *MethodMatch_ParameterMatchSpecifier) GetExactMatch() string {
	if x, ok := x.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (x *MethodMatch_ParameterMatchSpecifier) GetRangeMatch() *v3.Int64Range {
	if x, ok := x.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		return x.RangeMatch
	}
	return nil
}

type isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier interface {
	isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier()
}

type MethodMatch_ParameterMatchSpecifier_ExactMatch struct {
	// If specified, header match will be performed based on the value of the header.
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type MethodMatch_ParameterMatchSpecifier_RangeMatch struct {
	// If specified, header match will be performed based on range.
	// The rule will match if the request header value is within this range.
	// The entire request header value must represent an integer in base 10 notation: consisting
	// of an optional plus or minus sign followed by a sequence of digits. The rule will not match
	// if the header value does not represent an integer. Match will fail for empty values,
	// floating point numbers or if only a subsequence of the header value is an integer.
	//
	// Examples:
	//
	// * For range [-10,0), route will match for header value -1, but not for 0,
	//   "somestring", 10.9, "-1somestring"
	RangeMatch *v3.Int64Range `protobuf:"bytes,4,opt,name=range_match,json=rangeMatch,proto3,oneof"`
}

func (*MethodMatch_ParameterMatchSpecifier_ExactMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (*MethodMatch_ParameterMatchSpecifier_RangeMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

var File_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x12, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x52, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x48, 0x9a, 0xc5, 0x88, 0x1e, 0x43, 0x0a, 0x41, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a,
	0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x87, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x5f, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x60, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x3a, 0x3b, 0x9a, 0xc5,
	0x88, 0x1e, 0x36, 0x0a, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x22, 0xed, 0x01, 0x0a, 0x0a, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x58, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x43, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x40, 0x9a, 0xc5, 0x88, 0x1e, 0x3b, 0x0a, 0x39,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x22, 0xe2, 0x01, 0x0a, 0x0b, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x11, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65,
	0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x10, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x3a, 0x41, 0x9a, 0xc5, 0x88, 0x1e, 0x3c, 0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x18, 0x0a, 0x11, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x95,
	0x05, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3d,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x74, 0x0a,
	0x0c, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x1a, 0xf2, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x61, 0x63, 0x74, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x3c, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x3a, 0x59, 0x9a, 0xc5, 0x88, 0x1e, 0x54, 0x0a, 0x52, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x1b, 0x0a, 0x19, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x98, 0x01, 0x0a, 0x10, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x6e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x58,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x34, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x3a, 0x41, 0x9a, 0xc5, 0x88, 0x1e, 0x3c, 0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x59, 0x0a, 0x41, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6a, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0a, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescData = file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDesc
)

func file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDescData
}

var file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_goTypes = []interface{}{
	(*RouteConfiguration)(nil),                  // 0: envoy.extensions.filters.network.jres_proxy.v4alpha.RouteConfiguration
	(*Route)(nil),                               // 1: envoy.extensions.filters.network.jres_proxy.v4alpha.Route
	(*RouteMatch)(nil),                          // 2: envoy.extensions.filters.network.jres_proxy.v4alpha.RouteMatch
	(*RouteAction)(nil),                         // 3: envoy.extensions.filters.network.jres_proxy.v4alpha.RouteAction
	(*MethodMatch)(nil),                         // 4: envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch
	(*MethodMatch_ParameterMatchSpecifier)(nil), // 5: envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch.ParameterMatchSpecifier
	nil,                             // 6: envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch.ParamsMatchEntry
	(*v4alpha.HeaderMatcher)(nil),   // 7: envoy.config.route.v4alpha.HeaderMatcher
	(*v4alpha.WeightedCluster)(nil), // 8: envoy.config.route.v4alpha.WeightedCluster
	(*v4alpha1.StringMatcher)(nil),  // 9: envoy.type.matcher.v4alpha.StringMatcher
	(*v3.Int64Range)(nil),           // 10: envoy.type.v3.Int64Range
}
var file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_depIdxs = []int32{
	1,  // 0: envoy.extensions.filters.network.jres_proxy.v4alpha.RouteConfiguration.routes:type_name -> envoy.extensions.filters.network.jres_proxy.v4alpha.Route
	2,  // 1: envoy.extensions.filters.network.jres_proxy.v4alpha.Route.match:type_name -> envoy.extensions.filters.network.jres_proxy.v4alpha.RouteMatch
	3,  // 2: envoy.extensions.filters.network.jres_proxy.v4alpha.Route.route:type_name -> envoy.extensions.filters.network.jres_proxy.v4alpha.RouteAction
	4,  // 3: envoy.extensions.filters.network.jres_proxy.v4alpha.RouteMatch.method:type_name -> envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch
	7,  // 4: envoy.extensions.filters.network.jres_proxy.v4alpha.RouteMatch.headers:type_name -> envoy.config.route.v4alpha.HeaderMatcher
	8,  // 5: envoy.extensions.filters.network.jres_proxy.v4alpha.RouteAction.weighted_clusters:type_name -> envoy.config.route.v4alpha.WeightedCluster
	9,  // 6: envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch.name:type_name -> envoy.type.matcher.v4alpha.StringMatcher
	6,  // 7: envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch.params_match:type_name -> envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch.ParamsMatchEntry
	10, // 8: envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch.ParameterMatchSpecifier.range_match:type_name -> envoy.type.v3.Int64Range
	5,  // 9: envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch.ParamsMatchEntry.value:type_name -> envoy.extensions.filters.network.jres_proxy.v4alpha.MethodMatch.ParameterMatchSpecifier
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_init() }
func file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_init() {
	if File_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteConfiguration); i {
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
		file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
		file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatch); i {
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
		file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteAction); i {
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
		file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodMatch); i {
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
		file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodMatch_ParameterMatchSpecifier); i {
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
	file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
	file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*MethodMatch_ParameterMatchSpecifier_ExactMatch)(nil),
		(*MethodMatch_ParameterMatchSpecifier_RangeMatch)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto = out.File
	file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_rawDesc = nil
	file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_goTypes = nil
	file_envoy_extensions_filters_network_jres_proxy_v4alpha_route_proto_depIdxs = nil
}
