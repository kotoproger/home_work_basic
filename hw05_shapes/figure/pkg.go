package figure

import "errors"

type Shape interface {
	Area() float64
}

func CalculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("Переданный объект не фигура или фигура без площади")
	}

	return shape.Area(), nil
}
