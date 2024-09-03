package figure

import "errors"

type Shape interface {
	Area() (float64, error)
}

func CalculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не фигура или фигура без площади")
	}

	area, areaError := shape.Area()

	if areaError != nil {
		return 0, areaError
	}

	return area, nil
}
