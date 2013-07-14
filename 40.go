package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	values := strings.Fields(s)
	m := make(map[string]int)

	for _, v := range values {
		m[v] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
