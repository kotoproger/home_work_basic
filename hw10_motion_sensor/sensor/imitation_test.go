package sensor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadSensor(t *testing.T) {
	testCases := []struct {
		name           string
		durationMs     int
		freqMS         int
		expectedLength int
	}{
		{name: "10/3 -> 3", durationMs: 10, freqMS: 3, expectedLength: 3},
		{name: "3/10 -> 0", durationMs: 3, freqMS: 10, expectedLength: 0},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				ch := make(chan int, 100)
				duration := time.After(time.Millisecond * time.Duration(testCase.durationMs))
				freq := time.NewTicker(time.Millisecond * time.Duration(testCase.freqMS))
				ReadSensor(ch, duration, freq)
				assert.Equal(t, testCase.expectedLength, len(ch))
			},
		)
	}
}
