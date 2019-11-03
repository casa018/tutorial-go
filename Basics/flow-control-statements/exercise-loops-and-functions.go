package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SqrtDiff(x float64) float64 {
	z := 1.0
	lim := math.Pow10(-15)
	zdiff := 0.0
	for i := 0; ; i++ {
		zdiff = (z*z - x) / (2 * z)
		if math.Abs(zdiff) < lim {
			fmt.Println("loops: ", i)
			return z
		}
		z -= zdiff
	}
}

func main() {
	fmt.Println(Sqrt(10), math.Sqrt(10))
	fmt.Println(SqrtDiff(34), math.Sqrt(34))
}
