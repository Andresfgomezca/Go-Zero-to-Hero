package main

import (
	"fmt"
	"practice/animal"

)
func main(){
	a := animal.animal{
		name: "cat",
		age: "3",
		raze: "Feline",
		}
	fmt.Println(animal.animalDescription(a))
}