package main

import (
	"fmt"
	"strings"
)

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
	fmt.Println("----------------------------------------------")
	fmt.Println("Slices of slices")
	// Create a tic-tac-toe board. with slices inside slices
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	//this cicle is ready each row and separating the elements with  " "
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println("----------------------------------------------")
	fmt.Println("Appending to a slice")
	var aa []int
	printSlice(aa)
	//this function will add elements to a slice
	// append works on nil slices.
	aa = append(aa, 0)
	printSlice(aa)

	// The slice grows as needed.
	aa = append(aa, 1)
	printSlice(aa)

	// We can add more than one element at a time.
	aa = append(aa, 2, 3, 4)
	printSlice(aa)
	fmt.Println("----------------------------------------------")
	fmt.Println("Range")
	//range is a form of the for loop that will iterate over a slice and it will also return
	//2 elements, the index and the element
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println("----------------------------------------------")
	fmt.Println("Range Continued")
	//we can implement the for range with different variables if we want to omit a
	//value with _ instead of a variable
	pow = make([]int, 10)
	for i := range pow {
		// << means times 2 and unit(i) means how many times
		//in this case 2 will be repeated by 1 i times or 1x2xi where 1 can be modified
		//>> means divided by 2  and i will be how many times
		//10>>2=2.5
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	fmt.Println("----------------------------------------------")
	fmt.Println("Maps")
	var m map[string]Vertex
	//this map will map keys to values
	//where the keys are strings and the values are vertex
	m = make(map[string]Vertex)
	//here we are filling the map with a key called bell labs and it will contain a vertex
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Google"] = Vertex{37.42202, -122.08408}
	//now we are calling the vertex by looking the value with its key
	fmt.Println(m["Bell Labs"])
	//we can also print the complete map
	fmt.Println(m)
	fmt.Println("----------------------------------------------")
	fmt.Println("Mutating Maps")
	//these are the methods that mutate the maps
	ac := make(map[string]int)
	//inserting an element in the map
	ac["Answer"] = 42
	fmt.Println("The value:", ac["Answer"])
	//updating and element in the maap
	ac["Answer"] = 48
	fmt.Println("The value:", ac["Answer"])
	//deleting an element in the map
	delete(ac, "Answer")
	fmt.Println("The value:", ac["Answer"])
	//testing if the element is present in the map with the key
	ad, ok := m["Answer"]
	//the element will be retrieve to ad and ok will be a bool that defines if the key exists
	fmt.Println("The value:", ad, "Present?", ok)
	//if the key is present ok will be true and if it is not it will be false
	//if ad or ok are not declared we can use a short declaration form
	//ad, ok := m["Answer"]
	fmt.Println("----------------------------------------------")
	fmt.Println("function values")
	//functions can be used as values in other functions arguments or a return value for a function
	//func compute(fn func(float64, float64) float64) float64 {	return fn(3, 4)}
	fmt.Println("----------------------------------------------")
	fmt.Println("function closures")
	//those are functions that returns another function, the one that is return is 
	//defined anonymously in the body of the main function, this process allows the internal function
	//to have access to the variables of the main function
	//func adder() func(int) int {	sum := 0	return func(x int) int {sum += x	return sum	}}
	//pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
		//	pos(i),
		//	neg(-2*i),
		)
	}


}

type Vertex struct {
	X, Y float64
}

//we can also define a map a fill it out of the functions
//var m = map[string]Vertex{	"Bell Labs": {40.68433, -74.39967},	"Google":    {37.42202, -122.08408},}

//function created to print slice with information of its capacity and length
func printSlice(z []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(z), cap(z), z)
}
func printSlicel(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
