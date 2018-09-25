package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// see: https://gist.github.com/tetsuok/2281812
func fibonacci() func() int {
	// one based:
	//c1, c2 := 0, 1
	//return func() int {
	//	c := c2
	//	c2 = c1 + c
	//	c1 = c
	//	return c
	//}

	// zero based:
	prev, next := -1, 1
	return func() int {
		prev, next = next, prev+next
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
