package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.11, -74.442},
	"Google":    {37.23, -122.23},
}

func main() {
	fmt.Println(m)
}
