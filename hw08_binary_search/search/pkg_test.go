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
	}{
		{"empty list", []int{}, 0, -1},
		{"out of list", []int{1, 2, 3, 4, 5, 6}, 0, -1},
		{"not in list 2", []int{1, 2, 3, 5, 6, 7}, 4, -1},
		{"first", []int{1, 2, 3, 5, 6, 7}, 1, 0},
		{"last", []int{1, 2, 3, 5, 6, 7}, 7, 5},
		{"middle", []int{1, 2, 3, 5, 6, 7}, 3, 2},
	}

	for _, testCase := range TestCases {
		t.Run(testCase.name, func(t *testing.T) {
			index := BinarySearch(testCase.list, testCase.search)
			assert.Equal(t, testCase.foundedIndex, index)
		})
	}
}
