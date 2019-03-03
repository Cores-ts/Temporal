// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

package test

import (
	fmt "fmt"
	proto "gx/ipfs/QmddjPSGZb3ieihSseFeCfVRpZzcqczPNsD2DvarSwnjJB/gogo-protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Foo struct {
	A                    *int32   `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B                    *int32   `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	C                    []int32  `protobuf:"varint,3,rep,name=c" json:"c,omitempty"`
	D                    *int32   `protobuf:"varint,4,opt,name=d" json:"d,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetA() int32 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *Foo) GetB() int32 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

func (m *Foo) GetC() []int32 {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *Foo) GetD() int32 {
	if m != nil && m.D != nil {
		return *m.D
	}
	return 0
}

type Bar struct {
	Foos                 []*Foo   `protobuf:"bytes,1,rep,name=foos" json:"foos,omitempty"`
	Strs                 []string `protobuf:"bytes,2,rep,name=strs" json:"strs,omitempty"`
	Bufs                 [][]byte `protobuf:"bytes,3,rep,name=bufs" json:"bufs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}
func (m *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(m, src)
}
func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetFoos() []*Foo {
	if m != nil {
		return m.Foos
	}
	return nil
}

func (m *Bar) GetStrs() []string {
	if m != nil {
		return m.Strs
	}
	return nil
}

func (m *Bar) GetBufs() [][]byte {
	if m != nil {
		return m.Bufs
	}
	return nil
}

func init() {
	proto.RegisterType((*Foo)(nil), "test.Foo")
	proto.RegisterType((*Bar)(nil), "test.Bar")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0xcd, 0x21, 0x0e, 0xc3, 0x30,
	0x0c, 0x85, 0x61, 0xa5, 0xce, 0x40, 0xbd, 0xa2, 0x20, 0x93, 0x49, 0x51, 0x51, 0x50, 0xc1, 0xf8,
	0xc8, 0x40, 0xd1, 0x50, 0x6e, 0x90, 0xb4, 0x2b, 0xf5, 0x14, 0x67, 0xf7, 0xaf, 0xec, 0xb2, 0xff,
	0x7b, 0x91, 0x62, 0xc4, 0xfe, 0x95, 0xbe, 0xfc, 0x1a, 0x77, 0x0e, 0x5e, 0x7b, 0x7e, 0x21, 0xac,
	0xcc, 0x61, 0x42, 0x57, 0xc8, 0x45, 0x97, 0x6e, 0xd9, 0x15, 0x55, 0xa5, 0xe1, 0x52, 0x55, 0x6d,
	0x04, 0x11, 0x54, 0x9b, 0x6a, 0x27, 0x7f, 0xbd, 0xed, 0xf3, 0x07, 0xe1, 0x5d, 0x5a, 0x78, 0xa0,
	0x3f, 0x98, 0x85, 0x5c, 0x84, 0x74, 0x7f, 0x8e, 0x8b, 0x9d, 0x59, 0x99, 0xb3, 0xcd, 0x21, 0xa0,
	0x97, 0xde, 0x84, 0x86, 0x08, 0x69, 0xcc, 0xd6, 0xba, 0xd5, 0xff, 0x21, 0xf6, 0xf1, 0x94, 0xad,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xe7, 0x8e, 0xfd, 0x9f, 0x00, 0x00, 0x00,
}
