package sensor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDataProcessor(t *testing.T) {
	testCases := []struct {
		name         string
		input        []int
		expectedData []float32
	}{
		{name: "empty input", input: []int{}, expectedData: []float32{}},
		{name: "1 el", input: []int{1}, expectedData: []float32{}},
		{name: "2 el", input: []int{1, 2}, expectedData: []float32{}},
		{name: "3 el", input: []int{1, 2, 3}, expectedData: []float32{}},
		{name: "4 el", input: []int{1, 2, 3, 4}, expectedData: []float32{}},
		{name: "5 el", input: []int{1, 2, 3, 4, 5}, expectedData: []float32{}},
		{name: "6 el", input: []int{1, 2, 3, 4, 5, 6}, expectedData: []float32{}},
		{name: "7 el", input: []int{1, 2, 3, 4, 5, 6, 7}, expectedData: []float32{}},
		{name: "8 el", input: []int{1, 2, 3, 4, 5, 6, 7, 8}, expectedData: []float32{}},
		{name: "9 el", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expectedData: []float32{}},
		{name: "10 el", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expectedData: []float32{5.5}},
		{name: "11 el", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, expectedData: []float32{5.5}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outputCh := make(chan float32, len(testCase.expectedData)+100)
			defer close(outputCh)
			inputCh := make(chan int, len(testCase.input))
			defer close(inputCh)
			exit := time.After(time.Millisecond)
			for _, data := range testCase.input {
				inputCh <- data
			}
			DataProcessor(inputCh, outputCh, exit)
			actualData := make([]float32, len(outputCh))
			for i := 0; i < len(actualData); i++ {
				actualData[i] = <-outputCh
			}

			assert.Equal(t, testCase.expectedData, actualData)
		})
	}
}
