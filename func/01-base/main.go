package main

import (
	"fmt"
	"math"
)

func main() {
	x, _ := test(1, 2)

	fmt.Println(x)

	getSqrt := func(num float64) float64 {
		return math.Sqrt(num)
	}
	fmt.Println(getSqrt(4))
}

func test(x, y int) (int, int) {
	return x + 10, y + 10
}
