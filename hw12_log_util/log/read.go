package log

import (
	"bufio"
	"io"
)

func Read(reader io.Reader) <-chan []byte {
	output := make(chan []byte)

	go func(output chan<- []byte, reader io.Reader) {
		defer close(output)

		bReader := bufio.NewReader(reader)
		for {
			line, err := bReader.ReadString(byte('\n'))
			if err == io.EOF {
				return
			}
			output <- []byte(line)
		}
	}(output, reader)

	return output
}
