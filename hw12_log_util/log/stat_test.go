package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatCalc(t *testing.T) {
	testCases := []struct {
		name   string
		input  []Record
		output map[string]map[string]int
	}{
		{name: "empty chan", input: []Record{}, output: make(map[string]map[string]int)},
		{
			name: "some records",
			input: []Record{
				{Message: "123"},
				{},
				{Message: "123"},
				{Level: LogLevelFatal},
			},
			output: map[string]map[string]int{
				"Message": {
					"123": 2,
					"":    2,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				input := make(chan Record, len(testCase.input))
				for _, record := range testCase.input {
					input <- record
				}
				close(input)

				assert.Equal(t, testCase.output, StatCalc(input))
			},
		)
	}
}
