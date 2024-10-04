package book

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleCases = []struct {
	name string
	book Book
	json []byte
}{
	{"empty", Book{}, []byte(`{"idid":0,"title":"","author":"","year":0,"size":0,"rate":0}`)},
	{
		"not empty",
		Book{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6},
		[]byte(`{"idid":123,"title":"title","author":"author","year":2000,"size":100,"rate":1.6}`),
	},
}

var sliceCases = []struct {
	name  string
	books []Book
	bytes []byte
}{
	{"empty", []Book{}, []byte(`{"bOOks":[]}`)},
	{
		"one element",
		[]Book{{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6}},
		[]byte(`{"bOOks":[{"idid":123,"title":"title","author":"author","year":2000,"size":100,"rate":1.6}]}`),
	},
	{
		"two elements",
		[]Book{
			{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6},
			{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6},
		},
		[]byte(`{"bOOks":[{"idid":123,"title":"title","author":"author","year":2000,"size":100,"rate":1.6}` +
			`,{"idid":123,"title":"title","author":"author","year":2000,"size":100,"rate":1.6}]}`),
	},
}

func TestBookMarshalJSON(t *testing.T) {
	for _, tc := range simpleCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := json.Marshal(tc.book)
			assert.Equal(t, tc.json, res)
			assert.Nil(t, err)
		})
	}
}

func TestBookUnmarshalJSON(t *testing.T) {
	for _, tc := range simpleCases {
		t.Run(tc.name, func(t *testing.T) {
			b := Book{}
			err := b.UnmarshalJSON(tc.json)
			assert.Equal(t, tc.book, b)
			assert.Nil(t, err)
		})
	}

	t.Run("malformed json", func(t *testing.T) {
		b := Book{}
		err := b.UnmarshalJSON([]byte("sdfdfhawer"))
		assert.Equal(t, Book{}, b)
		assert.NotNil(t, err)
	})
}

func TestBooklistMarshalJson(t *testing.T) {
	for _, tc := range sliceCases {
		t.Run(tc.name, func(t *testing.T) {
			bl := Booklist{Books: make([]*Book, len(tc.books))}
			for index, book := range tc.books {
				book := book
				bl.Books[index] = &book
			}

			json, err := json.Marshal(bl)
			fmt.Print(string(json))
			assert.Nil(t, err)
			assert.Equal(t, tc.bytes, json)
		})
	}
}

func TestBooklistUnmarshalJSON(t *testing.T) {
	for _, tc := range sliceCases {
		t.Run(tc.name, func(t *testing.T) {
			bl := Booklist{Books: make([]*Book, len(tc.books))}
			for index, book := range tc.books {
				book := book
				bl.Books[index] = &book
			}

			result := Booklist{}
			err := json.Unmarshal(tc.bytes, &result)
			assert.Nil(t, err)
			assert.Equal(t, bl, result)
		})
	}
}
