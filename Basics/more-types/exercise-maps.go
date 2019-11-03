package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	var words []string = strings.Fields(s)
	for _, word := range words {
		m[word] = m[word] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
