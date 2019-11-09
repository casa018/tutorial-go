package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This methods means type T implements the interdace I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
