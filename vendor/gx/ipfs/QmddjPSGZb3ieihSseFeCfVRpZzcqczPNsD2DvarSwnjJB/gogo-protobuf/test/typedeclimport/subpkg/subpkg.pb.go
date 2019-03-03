// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subpkg/subpkg.proto

/*
Package subpkg is a generated protocol buffer package.

It is generated from these files:
	subpkg/subpkg.proto

It has these top-level messages:
	AnotherMessage
*/
package subpkg

import proto "gx/ipfs/QmddjPSGZb3ieihSseFeCfVRpZzcqczPNsD2DvarSwnjJB/gogo-protobuf/proto"
import fmt "fmt"
import math "math"
import _ "gx/ipfs/QmddjPSGZb3ieihSseFeCfVRpZzcqczPNsD2DvarSwnjJB/gogo-protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AnotherMessage struct {
	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (m *AnotherMessage) Reset()                    { *m = AnotherMessage{} }
func (m *AnotherMessage) String() string            { return proto.CompactTextString(m) }
func (*AnotherMessage) ProtoMessage()               {}
func (*AnotherMessage) Descriptor() ([]byte, []int) { return fileDescriptorSubpkg, []int{0} }

func (m *AnotherMessage) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func init() {
	proto.RegisterType((*AnotherMessage)(nil), "subpkg.AnotherMessage")
}
func (this *AnotherMessage) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AnotherMessage)
	if !ok {
		that2, ok := that.(AnotherMessage)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AnotherMessage")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AnotherMessage but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AnotherMessage but is not nil && this == nil")
	}
	if this.Foo != that1.Foo {
		return fmt.Errorf("Foo this(%v) Not Equal that(%v)", this.Foo, that1.Foo)
	}
	return nil
}
func (this *AnotherMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AnotherMessage)
	if !ok {
		that2, ok := that.(AnotherMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Foo != that1.Foo {
		return false
	}
	return true
}
func (m *AnotherMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnotherMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Foo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSubpkg(dAtA, i, uint64(len(m.Foo)))
		i += copy(dAtA[i:], m.Foo)
	}
	return i, nil
}

func encodeFixed64Subpkg(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Subpkg(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSubpkg(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AnotherMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovSubpkg(uint64(l))
	}
	return n
}

func sovSubpkg(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSubpkg(x uint64) (n int) {
	return sovSubpkg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnotherMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubpkg
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
			return fmt.Errorf("proto: AnotherMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnotherMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Foo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubpkg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSubpkg
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Foo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubpkg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSubpkg
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
func skipSubpkg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubpkg
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
					return 0, ErrIntOverflowSubpkg
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
					return 0, ErrIntOverflowSubpkg
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
				return 0, ErrInvalidLengthSubpkg
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSubpkg
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
				next, err := skipSubpkg(dAtA[start:])
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
	ErrInvalidLengthSubpkg = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubpkg   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("subpkg/subpkg.proto", fileDescriptorSubpkg) }

var fileDescriptorSubpkg = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0x4d, 0x2a,
	0xc8, 0x4e, 0xd7, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94,
	0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a, 0xbe,
	0x3e, 0x58, 0x3a, 0xa9, 0x34, 0x0d, 0xcc, 0x03, 0x73, 0xc0, 0x2c, 0x88, 0x36, 0x25, 0x25, 0x2e,
	0x3e, 0xc7, 0xbc, 0xfc, 0x92, 0x8c, 0xd4, 0x22, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21,
	0x01, 0x2e, 0xe6, 0xb4, 0xfc, 0x7c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0xd3, 0x49,
	0xe4, 0xc1, 0x43, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x24, 0x36, 0xb0, 0x01, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0x5d, 0xf6, 0x73, 0x8e, 0x00, 0x00, 0x00,
}
