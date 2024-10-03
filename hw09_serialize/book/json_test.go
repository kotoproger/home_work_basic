package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookMarshalJSON(t *testing.T) {
	testCases := []struct {
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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.book.MarshalJSON()
			assert.Equal(t, tc.json, res)
			assert.Nil(t, err)
		})
	}
}

func TestBookUnmarshalJSON(t *testing.T) {
	testCases := []struct {
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
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := Book{}
			err := b.UnmarshalJSON(tc.json)
			assert.Equal(t, tc.book, b)
			assert.Nil(t, err)
		})
	}
}
