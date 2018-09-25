package main

import (
	"fmt"
	"math"
)

// https://stackoverflow.com/questions/52441945/can-anyone-explain-the-computefn-func-code-in-the-go-tour-website
// func(float64, float64) float64 な関数を引数に持ち、float64を返却するcompute関数
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot)) // math.Sqrt(3*3 + 4*4): 5
	fmt.Println(math.Sqrt(3*3 + 4*4))
	fmt.Println(compute(math.Pow)) // math.Pow(3, 4)
	fmt.Println(math.Pow(3, 4))
}
