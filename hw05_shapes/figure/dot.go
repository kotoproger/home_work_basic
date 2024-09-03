package figure

type Dot struct {
	x int
	y int
}

func NewDot(x int, y int) *Dot {
	return &Dot{
		x: x,
		y: y,
	}
}
