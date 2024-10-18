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
		{name: "some random data", durationMs: 10, freqMS: 3},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actualData := []int{}
				duration := time.After(time.Millisecond * time.Duration(testCase.durationMs))
				ch := ReadSensor(duration, time.Millisecond*time.Duration(testCase.freqMS))
				for val := range ch {
					actualData = append(actualData, val)
				}
				assert.True(t, len(actualData) > 0)
			},
		)
	}
}
