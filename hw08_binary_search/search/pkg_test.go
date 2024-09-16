package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	TestCases := []struct {
		name         string
		list         []int
		search       int
		foundedIndex int
		founded      bool
	}{
		{"empty list", []int{}, 0, 0, false},
		{"out of list", []int{1, 2, 3, 4, 5, 6}, 0, 0, false},
		{"not in list 2", []int{1, 2, 3, 5, 6, 7}, 4, 0, false},
		{"first", []int{1, 2, 3, 5, 6, 7}, 1, 0, true},
		{"last", []int{1, 2, 3, 5, 6, 7}, 7, 5, true},
		{"middle", []int{1, 2, 3, 5, 6, 7}, 3, 2, true},
	}

	for _, testCase := range TestCases {
		t.Run(testCase.name, func(t *testing.T) {
			index, ok := BinarySearch(testCase.list, testCase.search)
			assert.Equal(t, testCase.foundedIndex, index)
			assert.Equal(t, testCase.founded, ok)
		})
	}
}
