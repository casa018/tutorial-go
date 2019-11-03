package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // golang don't use while statement. use for only.
		sum += sum
	}
	fmt.Println(sum)
}
