package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name   string
		input  []Record
		output []Record
		assert func(record Record) bool
	}{
		{name: "empty input", input: []Record{}, output: []Record{}, assert: func(_ Record) bool { return true }},
		{
			name:   "always true",
			input:  []Record{{Level: LogLevelDebug, Message: "1"}},
			output: []Record{{Level: LogLevelDebug, Message: "1"}},
			assert: func(_ Record) bool { return true },
		},
		{
			name:   "always false",
			input:  []Record{{Level: LogLevelDebug, Message: "1"}},
			output: []Record{},
			assert: func(_ Record) bool { return false },
		},
		{
			name: "some filter",
			input: []Record{
				{Level: LogLevelDebug, Message: "1"},
				{Level: LogLevelError, Message: "2"},
				{Level: LogLevelDebug, Message: "3"},
			},
			output: []Record{
				{Level: LogLevelDebug, Message: "1"},
				{Level: LogLevelDebug, Message: "3"},
			},
			assert: func(record Record) bool { return record.Level == LogLevelDebug },
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input := make(chan Record, len(testCase.input))
			for _, record := range testCase.input {
				input <- record
			}
			close(input)
			output := Filter(input, testCase.assert)
			outputSlice := []Record{}
			for outRecord := range output {
				outputSlice = append(outputSlice, outRecord)
			}

			assert.Equal(t, testCase.output, outputSlice)
		})
	}
}
