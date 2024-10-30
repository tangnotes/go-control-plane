// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.2
// source: envoy/extensions/common/aws/v3/credential_provider.proto

package awsv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Configuration for AWS credential provider. Normally, this is optional and the credentials are
// retrieved from the environment or AWS configuration files by following the default credential
// provider chain. This is to support cases where the credentials need to be explicitly provided
// by the control plane.
type AwsCredentialProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The option to use `AssumeRoleWithWebIdentity <https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRoleWithWebIdentity.html>`_.
	// If inline_credential is set, this is ignored.
	AssumeRoleWithWebIdentity *AssumeRoleWithWebIdentityCredentialProvider `protobuf:"bytes,1,opt,name=assume_role_with_web_identity,json=assumeRoleWithWebIdentity,proto3" json:"assume_role_with_web_identity,omitempty"`
	// The option to use an inline credential.
	// If this is set, it takes precedence over assume_role_with_web_identity.
	InlineCredential *InlineCredentialProvider `protobuf:"bytes,2,opt,name=inline_credential,json=inlineCredential,proto3" json:"inline_credential,omitempty"`
}

func (x *AwsCredentialProvider) Reset() {
	*x = AwsCredentialProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsCredentialProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsCredentialProvider) ProtoMessage() {}

func (x *AwsCredentialProvider) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsCredentialProvider.ProtoReflect.Descriptor instead.
func (*AwsCredentialProvider) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescGZIP(), []int{0}
}

func (x *AwsCredentialProvider) GetAssumeRoleWithWebIdentity() *AssumeRoleWithWebIdentityCredentialProvider {
	if x != nil {
		return x.AssumeRoleWithWebIdentity
	}
	return nil
}

func (x *AwsCredentialProvider) GetInlineCredential() *InlineCredentialProvider {
	if x != nil {
		return x.InlineCredential
	}
	return nil
}

// Configuration to use an inline AWS credential. This is an equivalent to setting the well-known
// environment variables “AWS_ACCESS_KEY_ID“, “AWS_SECRET_ACCESS_KEY“, and the optional “AWS_SESSION_TOKEN“.
type InlineCredentialProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The AWS access key ID.
	AccessKeyId string `protobuf:"bytes,1,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	// The AWS secret access key.
	SecretAccessKey string `protobuf:"bytes,2,opt,name=secret_access_key,json=secretAccessKey,proto3" json:"secret_access_key,omitempty"`
	// The AWS session token. This is optional.
	SessionToken string `protobuf:"bytes,3,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
}

func (x *InlineCredentialProvider) Reset() {
	*x = InlineCredentialProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InlineCredentialProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InlineCredentialProvider) ProtoMessage() {}

func (x *InlineCredentialProvider) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InlineCredentialProvider.ProtoReflect.Descriptor instead.
func (*InlineCredentialProvider) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescGZIP(), []int{1}
}

func (x *InlineCredentialProvider) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *InlineCredentialProvider) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

func (x *InlineCredentialProvider) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

// Configuration to use `AssumeRoleWithWebIdentity <https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRoleWithWebIdentity.html>`_
// to get AWS credentials.
type AssumeRoleWithWebIdentityCredentialProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ARN of the role to assume.
	RoleArn string `protobuf:"bytes,1,opt,name=role_arn,json=roleArn,proto3" json:"role_arn,omitempty"`
	// The web identity token that is provided by the identity provider to assume the role.
	WebIdentityToken string `protobuf:"bytes,2,opt,name=web_identity_token,json=webIdentityToken,proto3" json:"web_identity_token,omitempty"`
}

func (x *AssumeRoleWithWebIdentityCredentialProvider) Reset() {
	*x = AssumeRoleWithWebIdentityCredentialProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssumeRoleWithWebIdentityCredentialProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssumeRoleWithWebIdentityCredentialProvider) ProtoMessage() {}

func (x *AssumeRoleWithWebIdentityCredentialProvider) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssumeRoleWithWebIdentityCredentialProvider.ProtoReflect.Descriptor instead.
func (*AssumeRoleWithWebIdentityCredentialProvider) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescGZIP(), []int{2}
}

func (x *AssumeRoleWithWebIdentityCredentialProvider) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

func (x *AssumeRoleWithWebIdentityCredentialProvider) GetWebIdentityToken() string {
	if x != nil {
		return x.WebIdentityToken
	}
	return ""
}

var File_envoy_extensions_common_aws_v3_credential_provider_proto protoreflect.FileDescriptor

var file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x20, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x15, 0x41, 0x77, 0x73, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x8d,
	0x01, 0x0a, 0x1d, 0x61, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x77, 0x65, 0x62, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x57, 0x65, 0x62, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x19, 0x61, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x57, 0x65, 0x62, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x65,
	0x0a, 0x11, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x10, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0xaf, 0x01, 0x0a, 0x18, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x0d, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x2b, 0x41, 0x73, 0x73, 0x75,
	0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x57, 0x65, 0x62, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x3b, 0x0a, 0x12, 0x77,
	0x65, 0x62, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x52, 0x10, 0x77, 0x65, 0x62, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x9e, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x0a, 0x2c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76,
	0x33, 0x42, 0x17, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x77, 0x73,
	0x2f, 0x76, 0x33, 0x3b, 0x61, 0x77, 0x73, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescOnce sync.Once
	file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescData = file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDesc
)

func file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescGZIP() []byte {
	file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescData)
	})
	return file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDescData
}

var file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_common_aws_v3_credential_provider_proto_goTypes = []interface{}{
	(*AwsCredentialProvider)(nil),                       // 0: envoy.extensions.common.aws.v3.AwsCredentialProvider
	(*InlineCredentialProvider)(nil),                    // 1: envoy.extensions.common.aws.v3.InlineCredentialProvider
	(*AssumeRoleWithWebIdentityCredentialProvider)(nil), // 2: envoy.extensions.common.aws.v3.AssumeRoleWithWebIdentityCredentialProvider
}
var file_envoy_extensions_common_aws_v3_credential_provider_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.common.aws.v3.AwsCredentialProvider.assume_role_with_web_identity:type_name -> envoy.extensions.common.aws.v3.AssumeRoleWithWebIdentityCredentialProvider
	1, // 1: envoy.extensions.common.aws.v3.AwsCredentialProvider.inline_credential:type_name -> envoy.extensions.common.aws.v3.InlineCredentialProvider
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_common_aws_v3_credential_provider_proto_init() }
func file_envoy_extensions_common_aws_v3_credential_provider_proto_init() {
	if File_envoy_extensions_common_aws_v3_credential_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsCredentialProvider); i {
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
		file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InlineCredentialProvider); i {
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
		file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssumeRoleWithWebIdentityCredentialProvider); i {
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
			RawDescriptor: file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_common_aws_v3_credential_provider_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_common_aws_v3_credential_provider_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_common_aws_v3_credential_provider_proto_msgTypes,
	}.Build()
	File_envoy_extensions_common_aws_v3_credential_provider_proto = out.File
	file_envoy_extensions_common_aws_v3_credential_provider_proto_rawDesc = nil
	file_envoy_extensions_common_aws_v3_credential_provider_proto_goTypes = nil
	file_envoy_extensions_common_aws_v3_credential_provider_proto_depIdxs = nil
}