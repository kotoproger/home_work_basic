package figure

import "errors"

type Rectangle struct {
	length int
	width  int
}

func NewRectangle(length int, width int) *Rectangle {
	return &Rectangle{
		length: length,
		width:  width,
	}
}

func (rectangle *Rectangle) Area() (float64, error) {
	if rectangle.length > 0 && rectangle.width > 0 {
		return float64(rectangle.length * rectangle.width), nil
	}

	return 0, errors.New("не задана ширина или высота прямоугольника")
}
