package figure

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

func (rectangle *Rectangle) Area() float64 {
	return float64(rectangle.length * rectangle.width)
}
