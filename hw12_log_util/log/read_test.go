package log

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStringReader struct {
	buffer []byte
}

func (b *TestStringReader) Read(p []byte) (n int, err error) {
	for index := range p {
		if len(b.buffer) > 0 {
			n++
			p[index] = b.buffer[0]
			b.buffer = b.buffer[1:]
		} else {
			err = io.EOF
			p[index] = 0
		}
	}

	return
}

func TestRead(t *testing.T) {
	testCases := []struct {
		name   string
		input  []byte
		output [][]byte
	}{
		{name: "empty", input: []byte{}, output: [][]byte{}},
		{
			name:   "some strings",
			input:  []byte("asdasd\n\nsfsdf\njg563"),
			output: [][]byte{[]byte("asdasd\n"), []byte("\n"), []byte("sfsdf\n")},
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				output := Read(&TestStringReader{testCase.input})
				outputSlice := [][]byte{}
				for part := range output {
					outputSlice = append(outputSlice, part)
				}
				assert.Equal(
					t,
					testCase.output,
					outputSlice,
				)
			},
		)
	}
}
