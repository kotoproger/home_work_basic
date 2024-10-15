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
		closed       bool
	}{
		{name: "empty input", input: []int{}, expectedData: []float32{}, closed: false},
		{name: "1 el", input: []int{1}, expectedData: []float32{}, closed: false},
		{name: "2 el", input: []int{1, 2}, expectedData: []float32{}, closed: false},
		{name: "3 el", input: []int{1, 2, 3}, expectedData: []float32{}, closed: false},
		{name: "4 el", input: []int{1, 2, 3, 4}, expectedData: []float32{}, closed: false},
		{name: "5 el", input: []int{1, 2, 3, 4, 5}, expectedData: []float32{}, closed: false},
		{name: "6 el", input: []int{1, 2, 3, 4, 5, 6}, expectedData: []float32{}, closed: false},
		{name: "7 el", input: []int{1, 2, 3, 4, 5, 6, 7}, expectedData: []float32{}, closed: false},
		{name: "8 el", input: []int{1, 2, 3, 4, 5, 6, 7, 8}, expectedData: []float32{}, closed: false},
		{name: "9 el", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expectedData: []float32{}, closed: false},
		{name: "10 el", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expectedData: []float32{5.5}, closed: false},
		{name: "11 el", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, expectedData: []float32{5.5}, closed: false},
		{name: "empty input closed", input: []int{}, expectedData: []float32{}, closed: true},
		{name: "1 el closed", input: []int{1}, expectedData: []float32{}, closed: true},
		{name: "2 el closed", input: []int{1, 2}, expectedData: []float32{}, closed: true},
		{name: "3 el closed", input: []int{1, 2, 3}, expectedData: []float32{}, closed: true},
		{name: "4 el closed", input: []int{1, 2, 3, 4}, expectedData: []float32{}, closed: true},
		{name: "5 el closed", input: []int{1, 2, 3, 4, 5}, expectedData: []float32{}, closed: true},
		{name: "6 el closed", input: []int{1, 2, 3, 4, 5, 6}, expectedData: []float32{}, closed: true},
		{name: "7 el closed", input: []int{1, 2, 3, 4, 5, 6, 7}, expectedData: []float32{}, closed: true},
		{name: "8 el closed", input: []int{1, 2, 3, 4, 5, 6, 7, 8}, expectedData: []float32{}, closed: true},
		{name: "9 el closed", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expectedData: []float32{}, closed: true},
		{name: "10 el closed", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expectedData: []float32{5.5}, closed: true},
		{name: "11 el closed", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, expectedData: []float32{5.5}, closed: true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outputCh := make(chan float32, len(testCase.expectedData)+100)
			defer close(outputCh)
			inputCh := make(chan int, len(testCase.input))
			exit := time.After(time.Millisecond)
			for _, data := range testCase.input {
				inputCh <- data
			}
			if testCase.closed {
				close(inputCh)
			} else {
				defer close(inputCh)
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
