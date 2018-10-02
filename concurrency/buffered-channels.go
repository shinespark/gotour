package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 第2引数: バッファの長さ
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
