package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	f := strings.Fields(s)

	m := make(map[string]int)
	for _, w := range f {
		_, isPresent := m[w]
		if isPresent == false {
			m[w] = 1
		} else {
			m[w] += 1
		}

	}
	return m
}

func main() {
	wc.Test(WordCount)
}
