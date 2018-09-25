package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := 1.0

	for i := 1; i <= 10; i++ {
		diff := (z*z - x) / (2 * z)

		if z == z-diff {
			return z
		}

		z -= diff
	}

	return z

}

func main() {
	fmt.Println(Sqrt(2))

	for x := 1; x <= 10; x++ {
		fmt.Println(Sqrt(float64(x)))
		fmt.Println(math.Sqrt(float64(x)))
	}
}
