package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------------------------------------------")
	fmt.Println("Type parameters ")
	//this exercise implements the type parameters before the function arguments to
	//add the constraint comparable, this allows the function to implement the operators
	//== =! on values of the type T
	// Index works on a slice of ints
	//the function Index will receive the slice of ints and verifies if the value 15 is in the slice
	//returning its position or -1 if the value is not in the slice
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	//the function index will receive the slice of strings abd will verifies if the value
	//hello is in the slice, returning its position or returning -1 if the value is not
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
	fmt.Println("----------------------------------------------")
	fmt.Println("Generic types ")
	//the example focuses on using a struct that is form by a List of any type
	fmt.Println("----------------------------------------------")

}

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Index returns the index of x in s, or -1 if not found.
func Index[k comparable](s []k, x k) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}
