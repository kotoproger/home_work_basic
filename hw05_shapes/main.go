package main

import (
	"fmt"

	"github.com/kotoproger/home_work_basic/hw05_shapes/figure"
)

func main() {
	printSq(figure.NewCirlce(5))

	printSq(figure.NewRectangle(5, 10))

	printSq(figure.NewTriangle(8, 6))

	printSq(figure.NewDot(1, 2))

	printSq(figure.NewCirlce(0))

	printSq(figure.NewTriangle(8, 0))

	printSq(figure.NewRectangle(5, 0))
}

func printSq(s any) {
	square, calcError := figure.CalculateArea(s)
	if calcError == nil {
		fmt.Println(square)
	} else {
		fmt.Println(calcError)
	}
}
