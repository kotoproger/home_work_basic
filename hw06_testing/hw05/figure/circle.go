package figure

import (
	"errors"
	"math"
)

type Cirlce struct {
	radius int
}

func NewCirlce(radius int) *Cirlce {
	return &Cirlce{
		radius: radius,
	}
}

func (circle *Cirlce) Area() (float64, error) {
	if circle.radius > 0 {
		return math.Pi * float64(circle.radius) * float64(circle.radius), nil
	}

	return 0, errors.New("не задан радиус окружности")
}
