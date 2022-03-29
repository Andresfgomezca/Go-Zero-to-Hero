package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func main() {
	//this seed will now iterate
	rand.Seed(time.Now().UnixNano())
	fmt.Println("random number printed using rand and changing the seed to allow different ",
		"responses of the function: ", rand.Intn(10))
	//%g is used to locate the number that we are inserting with the function sqrt of math
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	//this part of code shows error if we do not call "pi" with capital letters,
	//it will not be exported in a different way
	fmt.Println(math.Pi)
	//this function will add 2 numbers and change the type of the inputs from int to float32
	fmt.Println(add(2, 1))
	// the swap function is returning 2 strings in the variables a and b
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	//in this case we are implenting the function split in 2 different forms, the first one is
	//just using the output submitting the input and the second one has 2 variables that will
	//receive the output of the function
	fmt.Println(split(31))
	x, y := split(31)
	fmt.Println(x, y)
	//the variable i is implementing the var statement to be part of a list of variables,
	//they are also initialized defining the initializer after the variables,
	//the initializer starts after the equals sign providing the value of the variables 1 by 1
	var i, j int = 1, 2
	fmt.Println(i, c, python, java, j)

	//the var statement can be ommited inside function if we use the short assignment statement :=
	//this option is not available at package level because the statements begins with keywords
	s, d, f := true, false, "no!"

	fmt.Println(s, d, f)
	//this section is showing the type of the variables and the respective value of the variable
	//using %t and %v
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

//this function will add 2 numbers, it is required to define the type of the variable after the
//variable

//implicit variable conversion is using variable type(variable)
func add(x int, y int) float32 {
	return float32(x + y)
}

//we can also omit the variable if parameters of the function share the same type by changing
//the parameters to x, y int

//a function can return  any number of results, it just need to be defined in the function
func swap(x, y string) (string, string) {
	return y, x
}

//Named return values, the function has the output defined next to the parameters
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//the var statement defines a list of variables and it can also be used in a function level
//the variables that does not have a type defined are bools by default and initialized in false

var c, python, java bool

//this is a different implementation of the var statement to create a list of variables
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//Variables declared without initial value will be initialized with their zero value.

//Constants are declared with the keyword Const
const Pi = 3.14
const World = "世界"
const Truth = true
