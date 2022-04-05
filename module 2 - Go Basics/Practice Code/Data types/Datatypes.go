package main

import (
	"fmt"
	"math/cmplx"
)

var (
	//variable of type bool
	ToBe   bool       = false
	//variable of type int unsigned 
	MaxInt uint64     = 1<<64 - 1
	//variabale complex with 15 significant decimal digits
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}