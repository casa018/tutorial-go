package main

import "fmt"

const Pi = 3.14

// dekiru
var i int = 10
var j int = 12

// dekinai -> 暗黙的な宣言は関数内のみ
// i := 10

func main() {
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
