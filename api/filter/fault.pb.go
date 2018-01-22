// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/filter/fault.proto

/*
	Package filter is a generated protocol buffer package.

	It is generated from these files:
		api/filter/fault.proto

	It has these top-level messages:
		FaultDelay
*/
package filter

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"
import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FaultDelay_FaultDelayType int32

const (
	// Fixed delay (step function).
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}
var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}
func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorFault, []int{0, 0}
}

// Delay specification is used to inject latency into the
// HTTP/gRPC/Mongo/Redis operation or delay proxying of TCP connections.
type FaultDelay struct {
	// Delay type to use (fixed|exponential|..). Currently, only fixed delay (step function) is
	// supported.
	Type FaultDelay_FaultDelayType `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.api.v2.filter.FaultDelay_FaultDelayType" json:"type,omitempty"`
	// An integer between 0-100 indicating the percentage of operations/connection requests
	// on which the delay will be injected.
	Percent uint32 `protobuf:"varint,2,opt,name=percent,proto3" json:"percent,omitempty"`
	// Types that are valid to be assigned to FaultDelayType:
	//	*FaultDelay_FixedDelay
	FaultDelayType isFaultDelay_FaultDelayType `protobuf_oneof:"fault_delay_type"`
}

func (m *FaultDelay) Reset()                    { *m = FaultDelay{} }
func (m *FaultDelay) String() string            { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()               {}
func (*FaultDelay) Descriptor() ([]byte, []int) { return fileDescriptorFault, []int{0} }

type isFaultDelay_FaultDelayType interface {
	isFaultDelay_FaultDelayType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FaultDelay_FixedDelay struct {
	FixedDelay *time.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,oneof,stdduration"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelayType() {}

func (m *FaultDelay) GetFaultDelayType() isFaultDelay_FaultDelayType {
	if m != nil {
		return m.FaultDelayType
	}
	return nil
}

func (m *FaultDelay) GetType() FaultDelay_FaultDelayType {
	if m != nil {
		return m.Type
	}
	return FaultDelay_FIXED
}

func (m *FaultDelay) GetPercent() uint32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *FaultDelay) GetFixedDelay() *time.Duration {
	if x, ok := m.GetFaultDelayType().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaultDelay_OneofMarshaler, _FaultDelay_OneofUnmarshaler, _FaultDelay_OneofSizer, []interface{}{
		(*FaultDelay_FixedDelay)(nil),
	}
}

func _FaultDelay_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaultDelay)
	// fault_delay_type
	switch x := m.FaultDelayType.(type) {
	case *FaultDelay_FixedDelay:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		dAtA, err := github_com_gogo_protobuf_types.StdDurationMarshal(*x.FixedDelay)
		if err != nil {
			return err
		}
		if err := b.EncodeRawBytes(dAtA); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FaultDelay.FaultDelayType has unexpected type %T", x)
	}
	return nil
}

func _FaultDelay_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaultDelay)
	switch tag {
	case 3: // fault_delay_type.fixed_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		if err != nil {
			return true, err
		}
		c := new(time.Duration)
		if err2 := github_com_gogo_protobuf_types.StdDurationUnmarshal(c, x); err2 != nil {
			return true, err
		}
		m.FaultDelayType = &FaultDelay_FixedDelay{c}
		return true, err
	default:
		return false, nil
	}
}

func _FaultDelay_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaultDelay)
	// fault_delay_type
	switch x := m.FaultDelayType.(type) {
	case *FaultDelay_FixedDelay:
		s := github_com_gogo_protobuf_types.SizeOfStdDuration(*x.FixedDelay)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FaultDelay)(nil), "envoy.api.v2.filter.FaultDelay")
	proto.RegisterEnum("envoy.api.v2.filter.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
}
func (m *FaultDelay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultDelay) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Type))
	}
	if m.Percent != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Percent))
	}
	if m.FaultDelayType != nil {
		nn1, err := m.FaultDelayType.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *FaultDelay_FixedDelay) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.FixedDelay != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFault(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.FixedDelay)))
		n2, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.FixedDelay, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func encodeVarintFault(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FaultDelay) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovFault(uint64(m.Type))
	}
	if m.Percent != 0 {
		n += 1 + sovFault(uint64(m.Percent))
	}
	if m.FaultDelayType != nil {
		n += m.FaultDelayType.Size()
	}
	return n
}

func (m *FaultDelay_FixedDelay) Size() (n int) {
	var l int
	_ = l
	if m.FixedDelay != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.FixedDelay)
		n += 1 + l + sovFault(uint64(l))
	}
	return n
}

func sovFault(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFault(x uint64) (n int) {
	return sovFault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FaultDelay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FaultDelay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaultDelay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (FaultDelay_FaultDelayType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percent", wireType)
			}
			m.Percent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Percent |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedDelay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := new(time.Duration)
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(v, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.FaultDelayType = &FaultDelay_FixedDelay{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFault
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFault
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFault
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFault(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFault = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFault   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/filter/fault.proto", fileDescriptorFault) }

var fileDescriptorFault = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0x4f, 0x4b, 0x2c, 0xcd, 0x29, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x4e, 0xcd, 0x2b, 0xcb, 0xaf, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33,
	0xd2, 0x83, 0x28, 0x90, 0x92, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x2b, 0x49, 0x2a,
	0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0x68, 0x92, 0x12, 0x2f, 0x4b,
	0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0xa0, 0x12, 0x22, 0xe9, 0xf9, 0xe9, 0xf9,
	0x60, 0xa6, 0x3e, 0x88, 0x05, 0x11, 0x55, 0x6a, 0x67, 0xe2, 0xe2, 0x72, 0x03, 0xd9, 0xe9, 0x92,
	0x9a, 0x93, 0x58, 0x29, 0xe4, 0xc7, 0xc5, 0x52, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x67, 0xa4, 0xa7, 0x87, 0xc5, 0x05, 0x7a, 0x08, 0xe5, 0x48, 0xcc, 0x90, 0xca, 0x82, 0x54,
	0x27, 0xae, 0x5d, 0x2f, 0x0f, 0x30, 0xb3, 0x36, 0x31, 0x32, 0x09, 0x30, 0x06, 0x81, 0xcd, 0x11,
	0x52, 0xe6, 0x62, 0x2f, 0x48, 0x2d, 0x4a, 0x4e, 0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x75, 0xe2, 0x04, 0x29, 0x61, 0xd1, 0x62, 0x92, 0x48, 0x09, 0x82, 0xc9, 0x08, 0xf9, 0x70, 0x71,
	0xa7, 0x65, 0x56, 0xa4, 0xa6, 0xc4, 0xa7, 0x80, 0x4c, 0x92, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x36,
	0x92, 0xd4, 0x83, 0x78, 0x54, 0x0f, 0xe6, 0x51, 0x3d, 0x17, 0xa8, 0x47, 0x9d, 0xf8, 0x66, 0xdc,
	0x97, 0x67, 0x04, 0x5b, 0xb5, 0x8a, 0x91, 0x49, 0x8b, 0xc1, 0x83, 0x21, 0x88, 0x0b, 0xac, 0x1f,
	0xec, 0x10, 0x25, 0x69, 0x2e, 0x3e, 0x54, 0x67, 0x09, 0x71, 0x72, 0xb1, 0xba, 0x79, 0x46, 0xb8,
	0xba, 0x08, 0x30, 0x38, 0x49, 0x72, 0x09, 0x80, 0x43, 0x18, 0x62, 0x55, 0x3c, 0xd8, 0x8d, 0xac,
	0x3b, 0x5e, 0x1e, 0x60, 0x66, 0x74, 0x12, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0xa3, 0xd8, 0x20, 0x9e, 0x4d, 0x62, 0x03, 0x5b, 0x6f, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x6b, 0x74, 0xe7, 0x25, 0xa4, 0x01, 0x00, 0x00,
}