// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/ip_tagging/v3/ip_tagging.proto

package envoy_extensions_filters_http_ip_tagging_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IPTagging_RequestType int32

const (
	IPTagging_BOTH     IPTagging_RequestType = 0
	IPTagging_INTERNAL IPTagging_RequestType = 1
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

var IPTagging_RequestType_name = map[int32]string{
	0: "BOTH",
	1: "INTERNAL",
	2: "EXTERNAL",
}

var IPTagging_RequestType_value = map[string]int32{
	"BOTH":     0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
}

func (x IPTagging_RequestType) String() string {
	return proto.EnumName(IPTagging_RequestType_name, int32(x))
}

func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_53ea551b2829685a, []int{0, 0}
}

type IPTagging struct {
	RequestType          IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=envoy.extensions.filters.http.ip_tagging.v3.IPTagging_RequestType" json:"request_type,omitempty"`
	IpTags               []*IPTagging_IPTag    `protobuf:"bytes,4,rep,name=ip_tags,json=ipTags,proto3" json:"ip_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IPTagging) Reset()         { *m = IPTagging{} }
func (m *IPTagging) String() string { return proto.CompactTextString(m) }
func (*IPTagging) ProtoMessage()    {}
func (*IPTagging) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ea551b2829685a, []int{0}
}

func (m *IPTagging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging.Unmarshal(m, b)
}
func (m *IPTagging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging.Marshal(b, m, deterministic)
}
func (m *IPTagging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging.Merge(m, src)
}
func (m *IPTagging) XXX_Size() int {
	return xxx_messageInfo_IPTagging.Size(m)
}
func (m *IPTagging) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging proto.InternalMessageInfo

func (m *IPTagging) GetRequestType() IPTagging_RequestType {
	if m != nil {
		return m.RequestType
	}
	return IPTagging_BOTH
}

func (m *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if m != nil {
		return m.IpTags
	}
	return nil
}

type IPTagging_IPTag struct {
	IpTagName            string          `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName,proto3" json:"ip_tag_name,omitempty"`
	IpList               []*v3.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList,proto3" json:"ip_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IPTagging_IPTag) Reset()         { *m = IPTagging_IPTag{} }
func (m *IPTagging_IPTag) String() string { return proto.CompactTextString(m) }
func (*IPTagging_IPTag) ProtoMessage()    {}
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ea551b2829685a, []int{0, 0}
}

func (m *IPTagging_IPTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging_IPTag.Unmarshal(m, b)
}
func (m *IPTagging_IPTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging_IPTag.Marshal(b, m, deterministic)
}
func (m *IPTagging_IPTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging_IPTag.Merge(m, src)
}
func (m *IPTagging_IPTag) XXX_Size() int {
	return xxx_messageInfo_IPTagging_IPTag.Size(m)
}
func (m *IPTagging_IPTag) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging_IPTag.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging_IPTag proto.InternalMessageInfo

func (m *IPTagging_IPTag) GetIpTagName() string {
	if m != nil {
		return m.IpTagName
	}
	return ""
}

func (m *IPTagging_IPTag) GetIpList() []*v3.CidrRange {
	if m != nil {
		return m.IpList
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.http.ip_tagging.v3.IPTagging_RequestType", IPTagging_RequestType_name, IPTagging_RequestType_value)
	proto.RegisterType((*IPTagging)(nil), "envoy.extensions.filters.http.ip_tagging.v3.IPTagging")
	proto.RegisterType((*IPTagging_IPTag)(nil), "envoy.extensions.filters.http.ip_tagging.v3.IPTagging.IPTag")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/ip_tagging/v3/ip_tagging.proto", fileDescriptor_53ea551b2829685a)
}

var fileDescriptor_53ea551b2829685a = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x71, 0xee, 0x0f, 0xa9, 0x73, 0x3a, 0x55, 0x59, 0xa8, 0x32, 0x1c, 0x47, 0xa7, 0x93,
	0x90, 0x6c, 0x94, 0x48, 0x70, 0x9c, 0x8e, 0x81, 0xa0, 0x93, 0xa8, 0x74, 0x94, 0x2a, 0xca, 0xc0,
	0x16, 0x99, 0x8b, 0x1b, 0x2c, 0xb5, 0xb6, 0xb1, 0xdd, 0xa8, 0x59, 0x99, 0x98, 0x19, 0x99, 0xf8,
	0x32, 0x7c, 0x1d, 0x3e, 0x00, 0x13, 0xb2, 0x1d, 0xda, 0x02, 0x53, 0xd9, 0xfc, 0xda, 0xef, 0xf3,
	0xfc, 0x1e, 0xbf, 0x7a, 0xe1, 0x35, 0xe5, 0xad, 0xe8, 0x30, 0x5d, 0x1b, 0xca, 0x35, 0x13, 0x5c,
	0xe3, 0x39, 0x5b, 0x18, 0xaa, 0x34, 0xfe, 0x60, 0x8c, 0xc4, 0x4c, 0x56, 0x86, 0x34, 0x0d, 0xe3,
	0x0d, 0x6e, 0xb3, 0x9d, 0x0a, 0x49, 0x25, 0x8c, 0x88, 0x1f, 0x3b, 0x35, 0xda, 0xaa, 0x51, 0xaf,
	0x46, 0x56, 0x8d, 0x76, 0xfa, 0xdb, 0x2c, 0x19, 0x7b, 0xd4, 0x9d, 0xe0, 0x73, 0xd6, 0xe0, 0x3b,
	0xa1, 0xa8, 0xf5, 0x24, 0x75, 0xad, 0xa8, 0xd6, 0xde, 0x30, 0x79, 0xb4, 0xaa, 0x25, 0xc1, 0x84,
	0x73, 0x61, 0x88, 0x71, 0x71, 0x5a, 0xaa, 0xac, 0xf3, 0x86, 0x99, 0x3c, 0x68, 0xc9, 0x82, 0xd5,
	0xc4, 0x50, 0xfc, 0xfb, 0xe0, 0x1f, 0xc6, 0x3f, 0x0e, 0xe0, 0x60, 0x32, 0x2b, 0x3d, 0x30, 0x16,
	0xf0, 0x44, 0xd1, 0x8f, 0x2b, 0xaa, 0x4d, 0x65, 0x3a, 0x49, 0x47, 0xe0, 0x1c, 0x5c, 0x9c, 0xa6,
	0x39, 0xda, 0x23, 0x31, 0xda, 0xb8, 0xa1, 0xc2, 0x5b, 0x95, 0x9d, 0xa4, 0x79, 0xf8, 0x33, 0x3f,
	0xfa, 0x04, 0x82, 0x21, 0x28, 0x22, 0xb5, 0xbd, 0x8e, 0x2b, 0x78, 0xdf, 0xab, 0xf5, 0xe8, 0xf0,
	0xfc, 0xe0, 0x22, 0x4a, 0xaf, 0xff, 0x93, 0xe5, 0x4e, 0x8e, 0xf2, 0x05, 0x04, 0x21, 0x28, 0x8e,
	0x99, 0x2c, 0x49, 0xa3, 0x93, 0x6f, 0x00, 0x1e, 0xb9, 0xb7, 0xf8, 0x0c, 0x46, 0x5e, 0x5c, 0x71,
	0xb2, 0xf4, 0x5f, 0x1b, 0x14, 0x03, 0xd7, 0x36, 0x25, 0x4b, 0x1a, 0x5f, 0xba, 0x28, 0x0b, 0xa6,
	0xcd, 0x28, 0x70, 0x51, 0x1e, 0xf6, 0x51, 0xfc, 0xec, 0x91, 0x9d, 0xbd, 0x65, 0xbe, 0x62, 0xb5,
	0x2a, 0x08, 0x6f, 0xa8, 0x65, 0xdc, 0x32, 0x6d, 0xae, 0x5e, 0x7c, 0xfd, 0xfe, 0xf9, 0xec, 0x12,
	0x3e, 0xfd, 0xa3, 0xdd, 0xa7, 0xfe, 0x37, 0x74, 0xfa, 0x77, 0xe8, 0x71, 0x06, 0xa3, 0x9d, 0x49,
	0xc5, 0x21, 0x3c, 0xcc, 0xdf, 0x96, 0xaf, 0x87, 0xf7, 0xe2, 0x13, 0x18, 0x4e, 0xa6, 0xe5, 0x4d,
	0x31, 0x7d, 0x79, 0x3b, 0x04, 0xb6, 0xba, 0x79, 0xd7, 0x57, 0xc1, 0xd5, 0x33, 0xcb, 0x4c, 0xe1,
	0x93, 0x7d, 0x99, 0xf9, 0x1b, 0xf8, 0x9c, 0x09, 0xff, 0x33, 0xa9, 0xc4, 0xba, 0xdb, 0x67, 0xde,
	0xf9, 0xe9, 0x44, 0xf6, 0x3e, 0x33, 0xbb, 0x3d, 0x33, 0xf0, 0xfe, 0xd8, 0xad, 0x51, 0xf6, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0xe1, 0xaa, 0x0e, 0x92, 0x13, 0x03, 0x00, 0x00,
}