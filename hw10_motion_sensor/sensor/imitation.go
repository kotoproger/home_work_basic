package sensor

import (
	"crypto/rand"
	"math/big"
	"time"
)

func ReadSensor(rawData chan int, exit <-chan time.Time, freq *time.Ticker) {
	for {
		select {
		case <-exit:
			return
		case <-freq.C:
			number, _ := rand.Int(rand.Reader, big.NewInt(100))
			rawData <- int(number.Int64())
		}
	}
}
