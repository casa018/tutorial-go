package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "World") //暗黙的な型の定義
	fmt.Println(a, b)
}
