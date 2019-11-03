package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	for y := 0; y < dy; y++ {
		var s []uint8
		for x := 0; x < dx; x++ {
			s = append(s, uint8((x+y)/2))
		}
		result = append(result, s)
	}
	return result
}

func main() {
	pic.Show(Pic)
}
