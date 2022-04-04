package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("----------------------------------------------")
	fmt.Println("Methods")
	//go does not have classes but  we can define methods on types,
	//this method will receive a type argument
	//the receiver appears between the func keyword and the method name
	//func (v Vertex) Abs() // keyword - receiver - method name
	//Abs() received a vertex named v
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println("----------------------------------------------")
	fmt.Println("Methods are functions ")
	//a method is just a function with a receiver argument,
	//the function abs2 is expressed as a regular function with no changes in functionality
	//the implementation is different
	fmt.Println(Abs2(v))
	fmt.Println("----------------------------------------------")
	fmt.Println("Methods contined")
	//we can declare a method on non-struct types but this types must be defined in the same package
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs3())
	fmt.Println("----------------------------------------------")
	fmt.Println("Pointer receivers ")
	//v := Vertex{3, 4}
	//the reason to define the pointer is to change the value of the variable inside the function
	//this means that the changes calculated in scale function will modify the values of the vertex
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println("----------------------------------------------")
	fmt.Println("Pointers and functions ")
	//in this section is shown that the functions abs2 and scale2 written as functions but
	//with pointers in their arguments, we also need to define the pointer for the argument of
	//the scale functio
	//v := Vertex{3, 4}
	Scale2(&v, 10)
	fmt.Println(Abs2(v))
	fmt.Println("----------------------------------------------")
	fmt.Println("Methods and pointer indirection ")
	//the functions with a pointer argument must take a pointer when they are called
	v.Scale(2)
	Scale2(&v, 10)
	//methods with pointer receivers take either a value or a pointer as the receiver when
	//they are called
	v.Scale(2)
	p := &Vertex{4, 3}
	p.Scale(3)
	Scale2(p, 8)

	fmt.Println(v, p)
	fmt.Println("----------------------------------------------")
	fmt.Println("Choosing a value or pointer receiver ")
	//effecs of the pointer in the receivers of these methods.
	v2 := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.Abs3())
	v.Scale(5)
	//the method scale is changing the value on the vertex that it receives
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.Abs3())
	fmt.Println("----------------------------------------------")
	fmt.Println("Interfaces ")
	//An interface defines empthy methods
	//it can also be define as a data type to be use in different functions or to build collecctions
	//structures can implement those interfaces
	//the example shows that the abstract method defined in the interface is implemented 2 times with
	//different inputs and it shows the error that appears by using the methods with an invalid input
	//as data not accepted in the interface
	fmt.Println("----------------------------------------------")
	fmt.Println("Interfaces are implemented implicitly ")
	//defining the structure t and calling the method of the interface
	//T is the struct and I the interface
	//M is the function defined in the interface I and implemented by T
	var i2 I
	i2 = &T{"hello"}
	i2.M()
	fmt.Println("----------------------------------------------")
	fmt.Println("Interface values ")
	//the interface values are defined of a specific type (value, type)
	//we can call diferent methods with the same name depending on the input defined in each method
	//the example has 2 methods called M where one received a float, it is also implementing the
	//type key word to define a type (type F float64) and the other one receives a struct
	//type T struct {	S string} that contains a String
	var ii I

	ii = &T{"Hello"}
	describe(ii)
	//Method m that receives a struct T
	ii.M()
	ii = F(math.Pi)
	//function that shows a message for this interface
	describe(ii)
	//Method m that receives a type F float64
	ii.M()
	fmt.Println("----------------------------------------------")
	fmt.Println("Interface values with nil underlying values ")
	//implementation of the interface on i3
	var i3 I
	//var t of type T with pointer
	var t *T
	//implementation of the interface explicitly
	i3 = t
	//shows the information of the var t
	describe(i3)
	//implementation of the method M with nil
	i3.M()
	//modification of the type T
	i3 = &T{"hello"}
	//shows the information of the interface
	describe(i3)
	//implementation of the method M with string "hello"
	i3.M()
	fmt.Println("----------------------------------------------")
	fmt.Println("Nil interface values ")
	//if we call a method on a nil interface it will run a run time error because there is no type
	//inside the interface tuple to indicate which concrete method to call
	//this means that if the interface is not related to any type or struct, the call
	//of a method will be an error
	fmt.Println("----------------------------------------------")
	fmt.Println("The empty interface ")
	//the empty interface shows that it can be implemented with different types and there will not be
	//an error by changing the interface value type
	var i5 interface{}
	describe2(i5)
	//empty interface receives an int
	i5 = 42
	describe2(i5)
	//empty interface receives a string
	i5 = "hello"
	describe2(i5)

	fmt.Println("----------------------------------------------")
	fmt.Println("Type assertioons ")
	//these assertions provides access to an interface value's
	var i6 interface{} = "hello"
	//this assert that the interface i holds the type string and assign its value
	s := i6.(string)
	fmt.Println(s)
	//this assert assign its value and returns true because the interface value is a string
	s, ok := i6.(string)
	fmt.Println(s, ok)
	//this asserts assign a 0 to f2 because it is a string and returns a false to ok
	f2, ok := i6.(float64)
	fmt.Println(f2, ok)
	//the assert of the wrong type will run an error
	//f2 = i6.(float64) // error due to the type of the interface that is string
	//fmt.Println(f)
	fmt.Println("----------------------------------------------")
	fmt.Println("Type switches ")
	//this switch will compare the type of the imput, that is what the function do process with
	//v:=i.(type)
	do(21)
	do("hello")
	do(true)
	//the type switch has the ssame syntax as a type assertion but the specific type T is replaced
	//with the keyword type and the value of the variable will be hold with the interface value
	fmt.Println("----------------------------------------------")
	fmt.Println("Stringers ")
	//changing the interface string allow us to print in a different way
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	fmt.Println("----------------------------------------------")
	fmt.Println("Errors ")
	//functions often returns error values, this if statement verify the status of the error
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("----------------------------------------------")
	fmt.Println(" ")

}

//errors
//error struct
type MyError struct {
	When time.Time
	What string
}

//function that prints the error information
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

//function that defines the error
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

//METHOD that describes empty interface
func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//method m that can receive a string that could be nil without having a null pointer exception
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type I interface {
	M()
}
type F float64

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func (f F) M() {
	fmt.Println(f)
}

//function defined in the interface I, T struct implements the interface due to the call
//in the method, this just means that the struct T implements the method d, it is not necesary to
//express it explicitly

type T struct {
	S string
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	//return root sqrt of the vertex axes
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	//increase f times the vertex attributes
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale2(v *Vertex, f float64) {
	//scale written as function with a pointer in its arguments
	v.X = v.X * f
	v.Y = v.Y * f
}

//methods abs3 that receives a pointer
func (v *Vertex) Abs3() float64 {
	//return root sqrt of the vertex axes
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
