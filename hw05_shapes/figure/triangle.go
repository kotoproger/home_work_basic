package figure

import "errors"

type Triangle struct {
	base   int
	height int
}

func NewTriangle(base int, height int) *Triangle {
	return &Triangle{
		base:   base,
		height: height,
	}
}

func (triangle *Triangle) Area() (float64, error) {
	if triangle.base > 0 && triangle.height > 0 {
		return float64(triangle.base*triangle.height) / 2, nil
	}

	return 0, errors.New("не задана высота или основание треугольника")
}
