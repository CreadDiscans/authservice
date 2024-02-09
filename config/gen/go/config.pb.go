// Copyright 2024 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: config.proto

package _go

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	mock "github.com/tetrateio/authservice-go/config/gen/go/mock"
	oidc "github.com/tetrateio/authservice-go/config/gen/go/oidc"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specifies how a request can be matched to a filter chain.
type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the http header used to match against.
	// Required.
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The criteria by which to match.
	// Must be one of `prefix` or `equality`.
	// Required.
	//
	// Types that are assignable to Criteria:
	//
	//	*Match_Prefix
	//	*Match_Equality
	Criteria isMatch_Criteria `protobuf_oneof:"criteria"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Match) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (m *Match) GetCriteria() isMatch_Criteria {
	if m != nil {
		return m.Criteria
	}
	return nil
}

func (x *Match) GetPrefix() string {
	if x, ok := x.GetCriteria().(*Match_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *Match) GetEquality() string {
	if x, ok := x.GetCriteria().(*Match_Equality); ok {
		return x.Equality
	}
	return ""
}

type isMatch_Criteria interface {
	isMatch_Criteria()
}

type Match_Prefix struct {
	// The expected prefix. If the actual value of the header starts with this prefix,
	// then it will be considered a match.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type Match_Equality struct {
	// The expected value. If the actual value of the header exactly equals this value,
	// then it will be considered a match.
	Equality string `protobuf:"bytes,3,opt,name=equality,proto3,oneof"`
}

func (*Match_Prefix) isMatch_Criteria() {}

func (*Match_Equality) isMatch_Criteria() {}

// A filter configuration.
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of filter. Currently, the only valid types are `oidc`
	// and `mock`. Required.
	//
	// Types that are assignable to Type:
	//
	//	*Filter_Oidc
	//	*Filter_OidcOverride
	//	*Filter_Mock
	Type isFilter_Type `protobuf_oneof:"type"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (m *Filter) GetType() isFilter_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Filter) GetOidc() *oidc.OIDCConfig {
	if x, ok := x.GetType().(*Filter_Oidc); ok {
		return x.Oidc
	}
	return nil
}

func (x *Filter) GetOidcOverride() *oidc.OIDCConfig {
	if x, ok := x.GetType().(*Filter_OidcOverride); ok {
		return x.OidcOverride
	}
	return nil
}

func (x *Filter) GetMock() *mock.MockConfig {
	if x, ok := x.GetType().(*Filter_Mock); ok {
		return x.Mock
	}
	return nil
}

type isFilter_Type interface {
	isFilter_Type()
}

type Filter_Oidc struct {
	// An OpenID Connect filter configuration.
	Oidc *oidc.OIDCConfig `protobuf:"bytes,1,opt,name=oidc,proto3,oneof"`
}

type Filter_OidcOverride struct {
	// This value will be used when `default_oidc_config` exists.
	// It will override values of them. If that doesn't exist,
	// this configuration will be rejected.
	OidcOverride *oidc.OIDCConfig `protobuf:"bytes,2,opt,name=oidc_override,json=oidcOverride,proto3,oneof"`
}

type Filter_Mock struct {
	// Mock filter configuration for testing and letting
	// AuthService run even if no OIDC providers are configured.
	Mock *mock.MockConfig `protobuf:"bytes,3,opt,name=mock,proto3,oneof"`
}

func (*Filter_Oidc) isFilter_Type() {}

func (*Filter_OidcOverride) isFilter_Type() {}

func (*Filter_Mock) isFilter_Type() {}

// A chain of one or more filters that will sequentially process an HTTP request.
type FilterChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A user-defined identifier for the processing chain used in log messages.
	// Required.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A rule to determine whether an HTTP request should be processed by the filter chain.
	// If not defined, the filter chain will match every request.
	// Optional.
	Match *Match `protobuf:"bytes,2,opt,name=match,proto3" json:"match,omitempty"`
	// The configuration of one of more filters in the filter chain. When the filter chain
	// matches an incoming request, then this list of filters will be applied to the request
	// in the order that they are declared.
	// All filters are evaluated until one of them returns a non-OK response.
	// If all filters return OK, the envoy proxy is notified that the request may continue.
	// The first filter that returns a non-OK response causes the request to be rejected with
	// the filter's returned status and any remaining filters are skipped.
	// At least one `Filter` is required in this array.
	Filters []*Filter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *FilterChain) Reset() {
	*x = FilterChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterChain) ProtoMessage() {}

func (x *FilterChain) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterChain.ProtoReflect.Descriptor instead.
func (*FilterChain) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *FilterChain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FilterChain) GetMatch() *Match {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *FilterChain) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

// The top-level configuration object.
// For a simple example, see the [sample JSON in the bookinfo configmap template](https://github.com/istio-ecosystem/authservice/blob/master/bookinfo-example/config/authservice-configmap-template-for-authn-and-authz.yaml).
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each incoming http request is matched against the list of filters in the chain, in order,
	// until a matching filter is found. The first matching filter is then applied to the request.
	// After the first match is made, other filters in the chain are ignored.
	// Order of chain declaration is therefore important.
	// At least one `FilterChain` is required in this array.
	Chains []*FilterChain `protobuf:"bytes,1,rep,name=chains,proto3" json:"chains,omitempty"`
	// The IP address for the Authservice to listen for incoming requests to process.
	// Required.
	ListenAddress string `protobuf:"bytes,2,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address,omitempty"`
	// The TCP port for the Authservice to listen for incoming requests to process.
	// Required.
	ListenPort int32 `protobuf:"varint,3,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	// The verbosity of logs generated by the Authservice.
	// Must be one of `trace`, `debug`, `info', 'error' or 'critical'.
	// Required.
	LogLevel string `protobuf:"bytes,4,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// The number of threads in the thread pool to use for processing.
	// The main thread will be used for accepting connections, before sending them to the thread-pool
	// for processing. The total number of running threads, including the main thread, will be N+1.
	// Required.
	Threads uint32 `protobuf:"varint,5,opt,name=threads,proto3" json:"threads,omitempty"`
	// List of trigger rules to decide if the Authservice should be used to authenticate the
	// request. The Authservice authentication happens if any one of the rules matched.
	// If the list is not empty and none of the rules matched, the request will be allowed
	// to proceed without Authservice authentication.
	// The format and semantics of `trigger_rules` are the same as the `triggerRules` setting
	// on the Istio Authentication Policy
	// (see https://istio.io/docs/reference/config/security/istio.authentication.v1alpha1).
	// CAUTION: Be sure that your configured `OIDCConfig.callback` and `OIDCConfig.logout` paths
	// each satisfies at least one of the trigger rules, or else the Authservice will not be able to
	// intercept requests made to those paths to perform the appropriate login/logout behavior.
	// Optional. Leave this empty to always trigger authentication for all paths.
	TriggerRules []*TriggerRule `protobuf:"bytes,9,rep,name=trigger_rules,json=triggerRules,proto3" json:"trigger_rules,omitempty"`
	// Global configuration of OIDC. This value will be applied to all filter definition
	// when it defined as `oidc_override`.
	// Optional.
	DefaultOidcConfig *oidc.OIDCConfig `protobuf:"bytes,10,opt,name=default_oidc_config,json=defaultOidcConfig,proto3" json:"default_oidc_config,omitempty"`
	// If true will allow the the requests even no filter chain match is found. Default false.
	// Optional.
	AllowUnmatchedRequests bool `protobuf:"varint,11,opt,name=allow_unmatched_requests,json=allowUnmatchedRequests,proto3" json:"allow_unmatched_requests,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *Config) GetChains() []*FilterChain {
	if x != nil {
		return x.Chains
	}
	return nil
}

func (x *Config) GetListenAddress() string {
	if x != nil {
		return x.ListenAddress
	}
	return ""
}

func (x *Config) GetListenPort() int32 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *Config) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *Config) GetThreads() uint32 {
	if x != nil {
		return x.Threads
	}
	return 0
}

func (x *Config) GetTriggerRules() []*TriggerRule {
	if x != nil {
		return x.TriggerRules
	}
	return nil
}

func (x *Config) GetDefaultOidcConfig() *oidc.OIDCConfig {
	if x != nil {
		return x.DefaultOidcConfig
	}
	return nil
}

func (x *Config) GetAllowUnmatchedRequests() bool {
	if x != nil {
		return x.AllowUnmatchedRequests
	}
	return false
}

// Trigger rule to match against a request. The trigger rule is satisfied if
// and only if both rules, excluded_paths and include_paths are satisfied.
type TriggerRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of paths to be excluded from the request. The rule is satisfied if
	// request path does not match to any of the path in this list.
	// Optional.
	ExcludedPaths []*StringMatch `protobuf:"bytes,1,rep,name=excluded_paths,json=excludedPaths,proto3" json:"excluded_paths,omitempty"`
	// List of paths that the request must include. If the list is not empty, the
	// rule is satisfied if request path matches at least one of the path in the list.
	// If the list is empty, the rule is ignored, in other words the rule is always satisfied.
	// Optional.
	IncludedPaths []*StringMatch `protobuf:"bytes,2,rep,name=included_paths,json=includedPaths,proto3" json:"included_paths,omitempty"`
}

func (x *TriggerRule) Reset() {
	*x = TriggerRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerRule) ProtoMessage() {}

func (x *TriggerRule) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerRule.ProtoReflect.Descriptor instead.
func (*TriggerRule) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *TriggerRule) GetExcludedPaths() []*StringMatch {
	if x != nil {
		return x.ExcludedPaths
	}
	return nil
}

func (x *TriggerRule) GetIncludedPaths() []*StringMatch {
	if x != nil {
		return x.IncludedPaths
	}
	return nil
}

// Describes how to match a given string. Match is case-sensitive.
type StringMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MatchType:
	//
	//	*StringMatch_Exact
	//	*StringMatch_Prefix
	//	*StringMatch_Suffix
	//	*StringMatch_Regex
	MatchType isStringMatch_MatchType `protobuf_oneof:"match_type"`
}

func (x *StringMatch) Reset() {
	*x = StringMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMatch) ProtoMessage() {}

func (x *StringMatch) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMatch.ProtoReflect.Descriptor instead.
func (*StringMatch) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (m *StringMatch) GetMatchType() isStringMatch_MatchType {
	if m != nil {
		return m.MatchType
	}
	return nil
}

func (x *StringMatch) GetExact() string {
	if x, ok := x.GetMatchType().(*StringMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *StringMatch) GetPrefix() string {
	if x, ok := x.GetMatchType().(*StringMatch_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *StringMatch) GetSuffix() string {
	if x, ok := x.GetMatchType().(*StringMatch_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (x *StringMatch) GetRegex() string {
	if x, ok := x.GetMatchType().(*StringMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

type isStringMatch_MatchType interface {
	isStringMatch_MatchType()
}

type StringMatch_Exact struct {
	// exact string match.
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatch_Prefix struct {
	// prefix-based match.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatch_Suffix struct {
	// suffix-based match.
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatch_Regex struct {
	// ECMAscript style regex-based match as defined by [EDCA-262](http://en.cppreference.com/w/cpp/regex/ecmascript).
	// Example: "^/pets/(.*?)?"
	Regex string `protobuf:"bytes,4,opt,name=regex,proto3,oneof"`
}

func (*StringMatch_Exact) isStringMatch_MatchType() {}

func (*StringMatch_Prefix) isStringMatch_MatchType() {}

func (*StringMatch_Suffix) isStringMatch_MatchType() {}

func (*StringMatch_Regex) isStringMatch_MatchType() {}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6f, 0x69, 0x64,
	0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x83, 0x01, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x25, 0x0a, 0x08, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x08, 0x65, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x0f, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xd7, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x39, 0x0a, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x12, 0x4a, 0x0a,
	0x0d, 0x6f, 0x69, 0x64, 0x63, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x4f,
	0x49, 0x44, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x69, 0x64,
	0x63, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x6d, 0x6f, 0x63,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x04,
	0x6d, 0x6f, 0x63, 0x6b, 0x42, 0x0b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42,
	0x01, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f,
	0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x3e, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22,
	0xe8, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x06, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x2e, 0x0a,
	0x0e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x70, 0x01, 0x52, 0x0d,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a,
	0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x10, 0x80, 0x80, 0x04, 0x52, 0x0a, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x47, 0x0a, 0x09, 0x6c, 0x6f, 0x67,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xfa, 0x42,
	0x27, 0x72, 0x25, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x08,
	0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x21, 0x0a, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x07, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0c, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x13, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6f, 0x69, 0x64, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x69,
	0x64, 0x63, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4f, 0x69, 0x64, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x38, 0x0a, 0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x75, 0x6e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x6e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x0d, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x50, 0x61, 0x74,
	0x68, 0x73, 0x12, 0x46, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0d, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x73, 0x22, 0x7f, 0x0a, 0x0b, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63,
	0x74, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x06, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x42, 0x0c, 0x0a,
	0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xc1, 0x01, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x74, 0x72, 0x61, 0x74, 0x65, 0x69, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02,
	0x12, 0x41, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0xca, 0x02, 0x12, 0x41, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xe2, 0x02, 0x1e, 0x41, 0x75, 0x74, 0x68, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41, 0x75, 0x74, 0x68,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_config_proto_goTypes = []interface{}{
	(*Match)(nil),           // 0: authservice.config.Match
	(*Filter)(nil),          // 1: authservice.config.Filter
	(*FilterChain)(nil),     // 2: authservice.config.FilterChain
	(*Config)(nil),          // 3: authservice.config.Config
	(*TriggerRule)(nil),     // 4: authservice.config.TriggerRule
	(*StringMatch)(nil),     // 5: authservice.config.StringMatch
	(*oidc.OIDCConfig)(nil), // 6: authservice.config.oidc.OIDCConfig
	(*mock.MockConfig)(nil), // 7: authservice.config.mock.MockConfig
}
var file_config_proto_depIdxs = []int32{
	6,  // 0: authservice.config.Filter.oidc:type_name -> authservice.config.oidc.OIDCConfig
	6,  // 1: authservice.config.Filter.oidc_override:type_name -> authservice.config.oidc.OIDCConfig
	7,  // 2: authservice.config.Filter.mock:type_name -> authservice.config.mock.MockConfig
	0,  // 3: authservice.config.FilterChain.match:type_name -> authservice.config.Match
	1,  // 4: authservice.config.FilterChain.filters:type_name -> authservice.config.Filter
	2,  // 5: authservice.config.Config.chains:type_name -> authservice.config.FilterChain
	4,  // 6: authservice.config.Config.trigger_rules:type_name -> authservice.config.TriggerRule
	6,  // 7: authservice.config.Config.default_oidc_config:type_name -> authservice.config.oidc.OIDCConfig
	5,  // 8: authservice.config.TriggerRule.excluded_paths:type_name -> authservice.config.StringMatch
	5,  // 9: authservice.config.TriggerRule.included_paths:type_name -> authservice.config.StringMatch
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterChain); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerRule); i {
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
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMatch); i {
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
	file_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Match_Prefix)(nil),
		(*Match_Equality)(nil),
	}
	file_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Filter_Oidc)(nil),
		(*Filter_OidcOverride)(nil),
		(*Filter_Mock)(nil),
	}
	file_config_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*StringMatch_Exact)(nil),
		(*StringMatch_Prefix)(nil),
		(*StringMatch_Suffix)(nil),
		(*StringMatch_Regex)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
