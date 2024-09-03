package figure

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

func (triangle *Triangle) Area() float64 {
	return float64(triangle.base*triangle.height) / 2
}
