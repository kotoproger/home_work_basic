package log

func Filter(input <-chan Record, assert func(record Record) bool) <-chan Record {
	output := make(chan Record)

	go func(output chan<- Record) {
		defer close(output)

		for record := range input {
			if assert(record) {
				output <- record
			}
		}
	}(output)

	return output
}
