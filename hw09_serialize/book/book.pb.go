// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book.proto

package book

import (
	fmt "fmt"
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

type Book struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	Year                 uint32   `protobuf:"varint,4,opt,name=Year,proto3" json:"Year,omitempty"`
	Size                 uint32   `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	Rate                 float32  `protobuf:"fixed32,6,opt,name=Rate,proto3" json:"Rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Book) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Book) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Book) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func init() {
	proto.RegisterType((*Book)(nil), "library.book")
}

func init() {
	proto.RegisterFile("book.proto", fileDescriptor_1e89d0eaa98dc5d8)
}

var fileDescriptor_1e89d0eaa98dc5d8 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xca, 0xcf, 0xcf,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0xc9, 0x4c, 0x2a, 0x4a, 0x2c, 0xaa, 0x54,
	0x6a, 0x60, 0xe4, 0x62, 0x01, 0x89, 0x0b, 0xf1, 0x71, 0x31, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xf0, 0x06, 0x31, 0x79, 0xba, 0x08, 0x89, 0x70, 0xb1, 0x86, 0x64, 0x96, 0xe4, 0xa4, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x62, 0x5c, 0x6c, 0x8e, 0xa5, 0x25, 0x19,
	0xf9, 0x45, 0x12, 0xcc, 0x60, 0x61, 0x28, 0x4f, 0x48, 0x88, 0x8b, 0x25, 0x32, 0x35, 0xb1, 0x48,
	0x82, 0x05, 0xac, 0x1f, 0xcc, 0x06, 0x89, 0x05, 0x67, 0x56, 0xa5, 0x4a, 0xb0, 0x42, 0xc4, 0x40,
	0x6c, 0x90, 0x58, 0x50, 0x62, 0x49, 0xaa, 0x04, 0x9b, 0x02, 0xa3, 0x06, 0x53, 0x10, 0x98, 0x9d,
	0xc4, 0x06, 0x76, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x80, 0xaf, 0x63, 0xa0, 0x00,
	0x00, 0x00,
}
