package figure

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	testCases := []struct {
		name      string
		length    int
		width     int
		area      float64
		calcError error
	}{
		{"empty length", 0, 10, 0, errors.New("не задана ширина или высота прямоугольника")},
		{"length less than 0", -2, 10, 0, errors.New("не задана ширина или высота прямоугольника")},
		{"empty width", 10, 0, 0, errors.New("не задана ширина или высота прямоугольника")},
		{"width less than 0", 10, -10, 0, errors.New("не задана ширина или высота прямоугольника")},
		{"noraml case", 3, 10, 30, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rect := NewRectangle(tc.length, tc.width)
			area, err := rect.Area()

			assert.Equal(t, tc.area, area)
			assert.Equal(t, tc.calcError, err)
		})
	}
}
