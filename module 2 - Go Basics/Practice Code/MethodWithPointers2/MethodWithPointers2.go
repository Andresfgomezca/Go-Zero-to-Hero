package main

import (
	"fmt"
)

type MyInt int

func (i MyInt) cube() int {
	return int(i) * int(i) * int(i)
}

func main() {
	a := MyInt(5)
	//method cube is returning the value of the type MyInt because the receiver receives MyInt type
	fmt.Println(a.cube())
}
