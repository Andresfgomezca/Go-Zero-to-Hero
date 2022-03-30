package main

import (
	"fmt"
)

func main() {

	x, y := Sqrt(164.65)
	fmt.Println("The loop has a sqrt root of value ", x, " and the cicle was repeated ", y, " times")

}

func Sqrt(x float64) (float64, int) {
	var z float64 = 1
	//loop of # interactions
	for i := 0; i < 1000; i++ {
		//tolerance defined 1%
		if (z*z) >= (x*0.99) && (z*z) <= (x*1.01) {
			return z, i
		} else {
			z -= (z*z - x) / (2 * z)
		}
	}
	return z, 1000
}
