package figure

import "errors"

type Shape interface {
	Area() (float64, error)
}

func CalculateArea(s any) (area float64, areaError error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не фигура или фигура без площади")
	}

	area, areaError = shape.Area()

	return
}
