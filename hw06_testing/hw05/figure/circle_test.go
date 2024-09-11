package figure

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAre(t *testing.T) {
	testCases := []struct {
		name      string
		radius    int
		area      float64
		calcError error
	}{
		{"empty radius", 0, 0, errors.New("не задан радиус окружности")},
		{"radius less than zero", -2, 0, errors.New("не задан радиус окружности")},
		{"redius 1", 1, 3.141592653589793, nil},
		{"redius 10", 10, 314.1592653589793, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			circle := NewCirlce(tc.radius)
			area, err := circle.Area()
			assert.Equal(t, tc.area, area)
			assert.Equal(t, err, tc.calcError)
		})
	}
}
