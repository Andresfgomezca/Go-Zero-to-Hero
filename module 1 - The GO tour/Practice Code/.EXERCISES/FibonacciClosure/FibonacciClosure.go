package main

import "fmt"

func fibonacci() func() int {
	//start point
	z, o := 0, 1
	return func() int {
		//each call will add the last 2 values
		z, o = o, z+o
		return o - z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
