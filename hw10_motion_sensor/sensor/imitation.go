package sensor

import (
	"crypto/rand"
	"math/big"
	"time"
)

func ReadSensor(exit <-chan time.Time, freq time.Duration) <-chan int {
	rawData := make(chan int)
	go func(rawData chan<- int) {
		defer close(rawData)
		for {
			number, _ := rand.Int(rand.Reader, big.NewInt(100))
			select {
			case <-exit:
				return
			case rawData <- int(number.Int64()):
			}
			time.Sleep(freq)
		}
	}(rawData)

	return rawData
}
