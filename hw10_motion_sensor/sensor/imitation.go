package sensor

import (
	"crypto/rand"
	"math/big"
	"time"
)

func ReadSensor(rawData chan<- int, exit <-chan time.Time, freq *time.Ticker) {
	for {
		select {
		case _, stillOpen := <-exit:
			if !stillOpen {
				return
			}
			return
		case _, stillOpen := <-freq.C:
			if !stillOpen {
				return
			}
			number, _ := rand.Int(rand.Reader, big.NewInt(100))
			rawData <- int(number.Int64())
		}
	}
}
