package figure

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaa(t *testing.T) {
	testCases := []struct {
		name      string
		base      int
		height    int
		area      float64
		calcError error
	}{
		{"empty base", 0, 1, 0, errors.New("не задана высота или основание треугольника")},
		{"base less than 0", -5, 1, 0, errors.New("не задана высота или основание треугольника")},
		{"empty height", 1, 0, 0, errors.New("не задана высота или основание треугольника")},
		{"height less than 0", 5, -1, 0, errors.New("не задана высота или основание треугольника")},
		{"normal case", 1, 4, 2, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			triangle := NewTriangle(tc.base, tc.height)
			area, err := triangle.Area()

			assert.Equal(t, tc.area, area)
			assert.Equal(t, tc.calcError, err)
		})
	}
}
