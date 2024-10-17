package sensor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadSensor(t *testing.T) {
	testCases := []struct {
		name       string
		durationMs int
		freqMS     int64
	}{
		{name: "10/3 -> 3", durationMs: 10, freqMS: 3},
		{name: "3/10 -> 0", durationMs: 3, freqMS: 100},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				ch := make(chan int, 100)
				defer close(ch)
				duration := time.After(time.Millisecond * time.Duration(testCase.durationMs))
				ReadSensor(ch, duration, time.Millisecond*3)
				assert.True(t, len(ch) > 0)
			},
		)
	}
}
