package main

import "fmt"

func main() {
	// a := []int{2, 3, 5, 7,11,13} 配列リテラル
	q := []int{2, 3, 5, 7, 11, 13} // a と同じ配列を作成し、それを参照するスライスの作成
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
