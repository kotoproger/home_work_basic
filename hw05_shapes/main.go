package main

import (
	"fmt"

	"github.com/kotoproger/home_work_basic/hw05_shapes/figure"
)

func main() {
	cirle := figure.NewCirlce(5)
	printSq(cirle)

	rect := figure.NewRectangle(5, 10)
	printSq(rect)

	triangle := figure.NewTriangle(8, 6)
	printSq(triangle)

	dot := figure.NewDot(1, 2)
	printSq(dot)
}

func printSq(s any) {
	sq, error := figure.CalculateArea(s)
	fmt.Println(sq, error)
}
