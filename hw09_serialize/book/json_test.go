package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleCases = []struct {
	name string
	book Book
	json []byte
}{
	{"empty", Book{}, []byte(`{"id":0,"title":"","author":"","year":0,"size":0,"rate":0}`)},
	{
		"not empty",
		Book{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6},
		[]byte(`{"id":123,"title":"title","author":"author","year":2000,"size":100,"rate":1.6}`),
	},
}

var sliceCases = []struct {
	name       string
	bookSlice  []Book
	jsonString []byte
}{
	{"empty", []Book{}, []byte("[]")},
	{
		"one element",
		[]Book{{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6}},
		[]byte(`[{"id":123,"title":"title","author":"author","year":2000,"size":100,"rate":1.6}]`),
	},
	{
		"two elements",
		[]Book{
			{ID: 123, Title: "title", Author: "author", Year: 2000, Size: 100, Rate: 1.6},
			{},
		},
		[]byte(`[{"id":123,"title":"title","author":"author","year":2000,"size":100,"rate":1.6}` +
			`,{"id":0,"title":"","author":"","year":0,"size":0,"rate":0}]`),
	},
}

func TestBookMarshalJSON(t *testing.T) {
	for _, tc := range simpleCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.book.MarshalJSON()
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

func TestMarshalBookSlice(t *testing.T) {
	for _, tc := range sliceCases {
		t.Run(tc.name, func(t *testing.T) {
			actualString, err := MarshalBookSlice(tc.bookSlice)
			assert.Equal(t, tc.jsonString, actualString)
			assert.Nil(t, err)
		})
	}
}

func TestUnmarshalBookSliceJSON(t *testing.T) {
	for _, tc := range sliceCases {
		t.Run(tc.name, func(t *testing.T) {
			actualSlice, err := UnmarshalBookSliceJSON(tc.jsonString)
			assert.Equal(t, tc.bookSlice, actualSlice)
			assert.Nil(t, err)
		})
	}

	t.Run("malformed json", func(t *testing.T) {
		actualSlice, err := UnmarshalBookSliceJSON([]byte("adgsdf"))
		assert.Nil(t, actualSlice)
		assert.NotNil(t, err)
	})
}
