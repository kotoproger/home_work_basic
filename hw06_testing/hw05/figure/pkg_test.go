package figure

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockShape struct {
	area float64
	err  error
}

func (m *MockShape) Area() (float64, error) {
	return m.area, m.err
}

func TestCalculateArea(t *testing.T) {
	testCases := []struct {
		name    string
		obj     any
		area    float64
		calcErr error
	}{
		{"int", 1, 0, errors.New("переданный объект не фигура или фигура без площади")},
		{"bool", true, 0, errors.New("переданный объект не фигура или фигура без площади")},
		{"string", "asdasd", 0, errors.New("переданный объект не фигура или фигура без площади")},
		{"float", 1.0, 0, errors.New("переданный объект не фигура или фигура без площади")},
		{"nil", nil, 0, errors.New("переданный объект не фигура или фигура без площади")},
		{"shape with area", &MockShape{area: 1.5, err: nil}, 1.5, nil},
		{"shape with area and error", &MockShape{23, errors.New("some error")}, 23, errors.New("some error")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			area, err := CalculateArea(tc.obj)
			assert.Equal(t, tc.area, area)
			assert.Equal(t, tc.calcErr, err)
		})
	}
}
