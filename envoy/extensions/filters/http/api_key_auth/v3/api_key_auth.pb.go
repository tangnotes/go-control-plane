// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/extensions/filters/http/api_key_auth/v3/api_key_auth.proto

package api_key_authv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

// API Key HTTP authentication.
//
// Example:
//
// .. code-block:: yaml
//
//	authentication_header: "X-API-KEY"
//	keys:
//	  inline_string: |-
//	    clientID1:apiKey1
//	    clientID2:apiKey2
type APIKeyAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// keys used to authenticate the client.
	// It should be a map of clientID to apiKey.
	// The clientID serves solely for identification purposes and isn't used for authentication.
	Keys *v3.DataSource `protobuf:"bytes,1,opt,name=keys,proto3" json:"keys,omitempty"`
	// The header name to fetch the key.
	// If multiple values are present in the given header, the filter rejects the request.
	// Only one of authentication_header, authentication_query, or authentication_cookie should be set.
	AuthenticationHeader string `protobuf:"bytes,2,opt,name=authentication_header,json=authenticationHeader,proto3" json:"authentication_header,omitempty"`
	// The query parameter name to fetch the key.
	// Only one of authentication_header, authentication_query, or authentication_cookie should be set.
	AuthenticationQuery string `protobuf:"bytes,3,opt,name=authentication_query,json=authenticationQuery,proto3" json:"authentication_query,omitempty"`
	// The cookie name to fetch the key.
	// Only one of authentication_header, authentication_query, or authentication_cookie should be set.
	AuthenticationCookie string `protobuf:"bytes,4,opt,name=authentication_cookie,json=authenticationCookie,proto3" json:"authentication_cookie,omitempty"`
}

func (x *APIKeyAuth) Reset() {
	*x = APIKeyAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKeyAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKeyAuth) ProtoMessage() {}

func (x *APIKeyAuth) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKeyAuth.ProtoReflect.Descriptor instead.
func (*APIKeyAuth) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescGZIP(), []int{0}
}

func (x *APIKeyAuth) GetKeys() *v3.DataSource {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *APIKeyAuth) GetAuthenticationHeader() string {
	if x != nil {
		return x.AuthenticationHeader
	}
	return ""
}

func (x *APIKeyAuth) GetAuthenticationQuery() string {
	if x != nil {
		return x.AuthenticationQuery
	}
	return ""
}

func (x *APIKeyAuth) GetAuthenticationCookie() string {
	if x != nil {
		return x.AuthenticationCookie
	}
	return ""
}

var File_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x3c, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x06, 0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x52, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x12, 0x33, 0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x15, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x42, 0xc5,
	0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08,
	0x01, 0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61,
	0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x42, 0x0f,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x3b, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescData = file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDesc
)

func file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDescData
}

var file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_goTypes = []interface{}{
	(*APIKeyAuth)(nil),    // 0: envoy.extensions.filters.http.api_key_auth.v3.APIKeyAuth
	(*v3.DataSource)(nil), // 1: envoy.config.core.v3.DataSource
}
var file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.api_key_auth.v3.APIKeyAuth.keys:type_name -> envoy.config.core.v3.DataSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_init() }
func file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_init() {
	if File_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKeyAuth); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto = out.File
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_rawDesc = nil
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_goTypes = nil
	file_envoy_extensions_filters_http_api_key_auth_v3_api_key_auth_proto_depIdxs = nil
}
