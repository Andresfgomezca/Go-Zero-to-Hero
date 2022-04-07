package main

import (
	"fmt"
)

//struct rectangle
type rect struct {
	width, height int
}

//method that receives a rectangle and calculates the area
func (r *rect) area() int {
	return r.width * r.height
}

//method that receives a rectangle and calculates the perimeter
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
