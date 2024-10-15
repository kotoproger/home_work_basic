package sensor

import (
	"time"
)

func DataProcessor(rawData <-chan int, output chan<- float32, exit <-chan time.Time) {
	counter := 0
	buffer := [10]int{}

	for {
		select {
		case _, stillOpen := <-exit:
			if !stillOpen {
				return
			}
			return
		case val, stillOpen := <-rawData:
			if !stillOpen {
				return
			}
			buffer[counter] = val
			counter++
			if counter == 10 {
				output <- calcucateData(buffer)
				counter = 0
			}
		}
	}
}

func calcucateData(data [10]int) float32 {
	var sum int
	for _, value := range data {
		sum += value
	}

	return float32(sum) / 10
}
