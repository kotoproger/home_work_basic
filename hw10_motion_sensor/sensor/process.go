package sensor

func DataProcessor(rawData <-chan int) <-chan float32 {
	output := make(chan float32)
	go func(output chan<- float32) {
		counter := 0
		buffer := 0
		defer close(output)

		for val := range rawData {
			buffer += val
			counter++
			if counter == 10 {
				output <- (float32(buffer) / float32(counter))
				counter = 0
			}
		}
	}(output)

	return output
}
