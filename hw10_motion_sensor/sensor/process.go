package sensor

import (
	"time"
)

func DataProcessor(rawData <-chan int, output chan<- float32, exit <-chan time.Time) {
	counter := 0
	buffer := [10]int{}

	for {
		select {
		case <-exit:
			return
		case buffer[counter] = <-rawData:
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
