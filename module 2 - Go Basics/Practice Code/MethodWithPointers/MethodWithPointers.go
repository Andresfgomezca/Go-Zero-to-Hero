package main

import "fmt"

type Calculator struct {
	num          int
	square, cube int
}

func (c *Calculator) calculateSquare() {
	c.square = c.num * c.num
}

func (c Calculator) calculateCube() {
	c.cube = c.num * c.num * c.num
}

func main() {
	c := Calculator{num: 5}

	c.calculateSquare()
	c.calculateCube()
	fmt.Println("Square of Num:", c.square)
	//how the method is not using pointer in the receiver it is not changing the cube value 
	//of the struct
	fmt.Println("Cube of Num:", c.cube)
}
