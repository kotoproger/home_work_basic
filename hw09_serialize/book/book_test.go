package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
)

var simpleCasesProtobuf = []struct {
	name string
	book Book
	json []byte
}{
	{"empty", Book{}, []byte{}},
	{
		"not empty",
		Book{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6},
		[]byte{
			0x8, 0x7b, 0x12, 0x5, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x6, 0x61, 0x75, 0x74,
			0x68, 0x6f, 0x72, 0x20, 0xd0, 0xf, 0x28, 0x64, 0x35, 0xcd, 0xcc, 0xcc, 0x3f,
		},
	},
}

var sliceCasesProtobuf = []struct {
	name  string
	books []Book
	bytes []byte
}{
	{"empty", []Book{}, []byte{}},
	{
		"one element",
		[]Book{{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6}},
		[]byte{
			0xa, 0x1b, 0x8, 0x7b, 0x12, 0x5, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x6, 0x61,
			0x75, 0x74, 0x68, 0x6f, 0x72, 0x20, 0xd0, 0xf, 0x28, 0x64, 0x35, 0xcd, 0xcc, 0xcc, 0x3f,
		},
	},
	{
		"two elements",
		[]Book{
			{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6},
			{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6},
		},
		[]byte{
			0xa, 0x1b, 0x8, 0x7b, 0x12, 0x5, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x6, 0x61,
			0x75, 0x74, 0x68, 0x6f, 0x72, 0x20, 0xd0, 0xf, 0x28, 0x64, 0x35, 0xcd, 0xcc, 0xcc,
			0x3f, 0xa, 0x1b, 0x8, 0x7b, 0x12, 0x5, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x6, 0x61,
			0x75, 0x74, 0x68, 0x6f, 0x72, 0x20, 0xd0, 0xf, 0x28, 0x64, 0x35, 0xcd, 0xcc, 0xcc, 0x3f,
		},
	},
}

func TestBookToString(t *testing.T) {
	for _, tc := range simpleCasesProtobuf {
		t.Run(tc.name, func(t *testing.T) {
			bookv2 := protoadapt.MessageV2Of(&tc.book)
			protoBytes, err := proto.Marshal(bookv2)
			assert.Equal(t, tc.json, protoBytes)
			assert.Nil(t, err)
		})
	}
}

func TestBookFromString(t *testing.T) {
	for _, tc := range simpleCasesProtobuf {
		t.Run(tc.name, func(t *testing.T) {
			b := Book{}
			bv2 := protoadapt.MessageV2Of(&b)
			err := proto.Unmarshal(tc.json, bv2)
			assert.Equal(t, tc.book, b)
			assert.Nil(t, err)
		})
	}
}

func TestBooksToString(t *testing.T) {
	for _, tc := range sliceCasesProtobuf {
		books := Booklist{Books: make([]*Book, len(tc.books))}
		for index, b := range tc.books {
			b := b
			books.Books[index] = &b
		}

		t.Run(tc.name, func(t *testing.T) {
			booksv2 := protoadapt.MessageV2Of(&books)
			protoBytes, err := proto.Marshal(booksv2)
			assert.Equal(t, tc.bytes, protoBytes)
			assert.Nil(t, err)
		})
	}
}

func TestBooksFromString(t *testing.T) {
	for _, tc := range sliceCasesProtobuf {
		books := Booklist{}
		if len(tc.books) > 0 {
			books.Books = make([]*Book, len(tc.books))
			for index, b := range tc.books {
				b := b
				books.Books[index] = &b
			}
		}

		t.Run(tc.name, func(t *testing.T) {
			list := Booklist{Books: make([]*Book, 0)}
			listv2 := protoadapt.MessageV2Of(&list)
			err := proto.Unmarshal(tc.bytes, listv2)
			assert.Equal(t, books, list)
			assert.Nil(t, err)
		})
	}
}
