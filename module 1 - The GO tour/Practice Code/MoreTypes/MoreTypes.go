package main

import "fmt"

func main() {
	fmt.Println("Implementation of the pointers that calls the memory address of a value")
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println("----------------------------------------------")
	fmt.Println("struct is a collection of fields ")
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println("----------------------------------------------")
	fmt.Println("those fields are accessed using a dot ")
	v.X = 4          //modifies x to 4
	fmt.Println(v.X) //access to X
	fmt.Println("----------------------------------------------")
	fmt.Println("we can also use pointers to have access to the structs")
	k := &v   //pointer to v
	k.X = 1e9 //the modification in k will change the X field in the struct v due to the pointer
	fmt.Println(v)
	fmt.Println("----------------------------------------------")
	fmt.Println("struct literals denotes a newly allocated struct value by listing ",
		"the values of its fields")
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		l  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, l, v2, v3)
	fmt.Println("----------------------------------------------")
	fmt.Println("arrays are described by the type next to the size, can't be resized")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("Array of strings")
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println("Array of primes numbers")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println("----------------------------------------------")
	fmt.Println("Slices")
	//an slice is dynamically-sized, just need to be crerated with the type is formed by specifying
	//two indices the start point and the end but the last one will be excluded
	var s []int = primes[1:4] //starts with element 1 but the element 4 will be excluded
	fmt.Println(s)
	fmt.Println("----------------------------------------------")
	fmt.Println("Slices does not store any data, so the changes in the slice will be",
		" applied in the underlying array")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	r := names[0:2]
	e := names[1:3]
	fmt.Println(r, e)

	e[0] = "XXX" //this change will modify the array of names
	fmt.Println(r, e)
	fmt.Println(names)
	fmt.Println("----------------------------------------------")
	fmt.Println("slice literals")
	//this literals are created like the arrays but just without the size
	//literals are used to construct the values for arrays, structs, slices,
	//and maps. Each time they are evaluated
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	u := []bool{true, false, true, true, false, true}
	fmt.Println(u)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)
	fmt.Println("----------------------------------------------")
	fmt.Println("slice defaults")
	//slices has defauls values when they are not specified, this one is 0 for low bound
	//and complete length for high bound
	w := []int{2, 3, 5, 7, 11, 13}

	w = w[1:4]
	fmt.Println(w)

	w = w[:2]
	fmt.Println(w)

	w = w[1:]
	fmt.Println(w)
	fmt.Println("----------------------------------------------")
	fmt.Println("slice length and capacity")
	z := []int{2, 3, 5, 7, 11, 13}
	printSlice(z)
	//the length of the slice can not be higher than its capacity

	// Slice the slice to give it zero length.
	z = z[:0]
	printSlice(z)

	// Extend its length.
	z = z[:4]
	printSlice(z)

	// Drop its first two values.
	z = z[2:]
	printSlice(z)
	fmt.Println("----------------------------------------------")
	fmt.Println("nil slice")
	//a nil slice is one that has length and capacity equals to 0 this means no underlying array
	var n []int
	printSlice(n)
	if n == nil {
		fmt.Println("nil!")
	}
	fmt.Println("----------------------------------------------")
	fmt.Println("creating a slice with make")
	//to create a slice with the build in make function
	f := make([]int, 5)
	printSlicel("f", f)

	b := make([]int, 0, 5)
	printSlicel("b", b)

	c := b[:2]
	printSlicel("c", c)

	d := c[2:5]
	printSlicel("d", d)
}

type Vertex struct {
	X int
	Y int
}

//function created to print slice with information of its capacity and length
func printSlice(z []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(z), cap(z), z)
}
func printSlicel(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
