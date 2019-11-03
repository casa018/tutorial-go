package main

import "fmt"

func main() {
	i, j := 42, 2101

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(*p) // see the new value of i

	p = &j       // point to j
	*p = *p / 37 //devide l through the pointer
	fmt.Println(j)
}
