// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book/book.proto

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
	return fileDescriptor_ee9082fb44230b1b, []int{0}
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

type Booklist struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Booklist) Reset()         { *m = Booklist{} }
func (m *Booklist) String() string { return proto.CompactTextString(m) }
func (*Booklist) ProtoMessage()    {}
func (*Booklist) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee9082fb44230b1b, []int{1}
}

func (m *Booklist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Booklist.Unmarshal(m, b)
}
func (m *Booklist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Booklist.Marshal(b, m, deterministic)
}
func (m *Booklist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Booklist.Merge(m, src)
}
func (m *Booklist) XXX_Size() int {
	return xxx_messageInfo_Booklist.Size(m)
}
func (m *Booklist) XXX_DiscardUnknown() {
	xxx_messageInfo_Booklist.DiscardUnknown(m)
}

var xxx_messageInfo_Booklist proto.InternalMessageInfo

func (m *Booklist) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

func init() {
	proto.RegisterType((*Book)(nil), "book.book")
	proto.RegisterType((*Booklist)(nil), "book.booklist")
}

func init() {
	proto.RegisterFile("book/book.proto", fileDescriptor_ee9082fb44230b1b)
}

var fileDescriptor_ee9082fb44230b1b = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0x3f, 0x0b, 0xc2, 0x30,
	0x10, 0xc5, 0x49, 0xfa, 0x07, 0x3d, 0x51, 0xe1, 0x10, 0xb9, 0x31, 0x74, 0xca, 0x20, 0x15, 0xf4,
	0x13, 0x08, 0x5d, 0xba, 0x46, 0x17, 0xc7, 0x16, 0x02, 0x16, 0x0b, 0x91, 0x36, 0x2e, 0x4e, 0x7e,
	0x74, 0xb9, 0x8b, 0xcb, 0xe3, 0xf7, 0x7e, 0xc7, 0xc1, 0x83, 0x6d, 0x1f, 0xc2, 0xf3, 0xc8, 0x51,
	0xbf, 0xa6, 0x10, 0x03, 0xe6, 0xcc, 0xd5, 0x57, 0x81, 0x00, 0x6e, 0x40, 0xb7, 0x0d, 0x29, 0xa3,
	0xec, 0xda, 0xe9, 0xb6, 0xc1, 0x1d, 0x14, 0xb7, 0x21, 0x8e, 0x9e, 0xb4, 0x51, 0x76, 0xe9, 0x52,
	0xc1, 0x3d, 0x94, 0x97, 0x77, 0x7c, 0x84, 0x89, 0x32, 0xd1, 0xff, 0x86, 0x08, 0xf9, 0xdd, 0x77,
	0x13, 0xe5, 0xf2, 0x2f, 0xcc, 0xee, 0x3a, 0x7c, 0x3c, 0x15, 0xc9, 0x31, 0xb3, 0x73, 0x5d, 0xf4,
	0x54, 0x1a, 0x65, 0xb5, 0x13, 0xae, 0x0e, 0xb0, 0xe0, 0x05, 0xe3, 0x30, 0x47, 0x34, 0x50, 0x30,
	0xcf, 0xa4, 0x4c, 0x66, 0x57, 0x27, 0xa8, 0x65, 0x30, 0x87, 0x4b, 0x87, 0xbe, 0x94, 0xf5, 0xe7,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x60, 0x3a, 0x45, 0xd0, 0x00, 0x00, 0x00,
}
