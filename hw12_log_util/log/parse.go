package log

import "encoding/json"

func Parse(input <-chan []byte) <-chan Record {
	output := make(chan Record)

	go func(output chan<- Record) {
		defer close(output)

		for slice := range input {
			record := Record{}
			err := json.Unmarshal(
				slice,
				&record,
			)
			if err == nil {
				output <- record
			}
		}
	}(output)

	return output
}
