package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	i := 42
	f2 := math.Sqrt(float64(i))
	u := uint(f2)
	fmt.Println(i, f2, u)
}
