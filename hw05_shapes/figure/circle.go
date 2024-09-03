package figure

import "math"

type Cirlce struct {
	radius int
}

func NewCirlce(radius int) *Cirlce {
	return &Cirlce{
		radius: radius,
	}
}

func (circle *Cirlce) Area() float64 {
	return math.Pi * float64(circle.radius) * float64(circle.radius)
}
