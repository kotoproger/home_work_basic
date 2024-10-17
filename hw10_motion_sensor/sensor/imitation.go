package sensor

import (
	"crypto/rand"
	"math/big"
	"time"
)

func ReadSensor(rawData chan<- int, exit <-chan time.Time, freq time.Duration) {
	for {
		number, _ := rand.Int(rand.Reader, big.NewInt(100))
		select {
		case _, stillOpen := <-exit:
			if !stillOpen {
				return
			}
			return
		case rawData <- int(number.Int64()):
			time.Sleep(freq)
		default:
			time.Sleep(freq)
		}
	}
}
