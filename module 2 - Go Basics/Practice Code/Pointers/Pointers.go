package main

import (
	"fmt"
)

func main() {
	//POINTERS
	fmt.Println("Pointers")
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Println(*&a)
	//Updated the value of a Using the pointer
	*b = 43
	fmt.Println(a)

	//Pass By Reference Using Pointer
	fmt.Println("Pass By Reference Using Pointer")
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer:", &i)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
