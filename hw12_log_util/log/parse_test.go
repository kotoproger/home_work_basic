package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]byte
		output []Record
	}{
		{name: "empty input", input: [][]byte{}, output: []Record{}},
		{name: "malformed json", input: [][]byte{[]byte("{\"level\":")}, output: []Record{}},
		{name: "empty record", input: [][]byte{[]byte("{}")}, output: []Record{{}}},
		{name: "with extra fields", input: [][]byte{[]byte("{\"som_field\":234}")}, output: []Record{{}}},
		{
			name:   "some records",
			input:  [][]byte{[]byte("{}"), []byte("{\"level\":\"info\"}")},
			output: []Record{{}, {Level: LogLevelInfo}},
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				input := make(chan []byte, len(testCase.input))
				for _, part := range testCase.input {
					input <- part
				}
				close(input)

				output := Parse(input)

				outputSlice := []Record{}
				for record := range output {
					outputSlice = append(outputSlice, record)
				}

				assert.Equal(t, testCase.output, outputSlice)
			},
		)
	}
}
