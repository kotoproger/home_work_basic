package sensor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataProcessor(t *testing.T) {
	testCases := []struct {
		name         string
		input        []int
		expectedData []float32
	}{
		{name: "empty input closed", input: []int{}, expectedData: []float32{}},
		{name: "1 el closed", input: []int{1}, expectedData: []float32{}},
		{name: "2 el closed", input: []int{1, 2}, expectedData: []float32{}},
		{name: "3 el closed", input: []int{1, 2, 3}, expectedData: []float32{}},
		{name: "4 el closed", input: []int{1, 2, 3, 4}, expectedData: []float32{}},
		{name: "5 el closed", input: []int{1, 2, 3, 4, 5}, expectedData: []float32{}},
		{name: "6 el closed", input: []int{1, 2, 3, 4, 5, 6}, expectedData: []float32{}},
		{name: "7 el closed", input: []int{1, 2, 3, 4, 5, 6, 7}, expectedData: []float32{}},
		{name: "8 el closed", input: []int{1, 2, 3, 4, 5, 6, 7, 8}, expectedData: []float32{}},
		{name: "9 el closed", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expectedData: []float32{}},
		{name: "10 el closed", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expectedData: []float32{5.5}},
		{name: "11 el closed", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, expectedData: []float32{5.5}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			inputCh := make(chan int, len(testCase.input))
			for _, data := range testCase.input {
				inputCh <- data
			}
			close(inputCh)

			outputCh := DataProcessor(inputCh)
			actualData := make([]float32, len(testCase.expectedData))
			index := 0
			for value := range outputCh {
				actualData[index] = value
				index++
			}

			assert.Equal(t, testCase.expectedData, actualData)
		})
	}
}
