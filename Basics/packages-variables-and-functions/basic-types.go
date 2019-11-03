package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1            // 1を左に64bitずらして1を引く -> 64bit最大値(64桁1が並ぶ)
	z      complex128 = cmplx.Sqrt(-5 + 12i) // 複素数型
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
