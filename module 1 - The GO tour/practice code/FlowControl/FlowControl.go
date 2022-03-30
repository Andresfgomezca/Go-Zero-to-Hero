package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	sum1 := 0
	//the for loop, does not use parentheses but braces {} are always required
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println("for loop", sum1)
	fmt.Println("----------------------------------------------")
	// For is the Go's "while", why the response is not changing according to aa correct response???
	sum2 := 1
	for sum2 < 62 {
		sum2 += sum2
	}
	fmt.Println("while loop", sum2)
	fmt.Println("----------------------------------------------")
	//if statements works like the for loop, parentheses is not required but the braces {} are mandatory
	if sum1 > sum2 {
		sum1 = sum2
	}
	fmt.Println("----------------------------------------------")
	//implementation of the pow function with if statements, but the else of the function is executed before the println. Whyy???
	fmt.Println("WHY THIS RESPONSE IS AFTER THE POW FUNCTION?",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println("----------------------------------------------")
	//the switch statement has implemented the break at the end of each case and the values
	//involve dont need to be integers
	//the variable os is being defined in the scope of the switch statement.
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	fmt.Println("----------------------------------------------")
	//another switch evaluation but this time it is running with time
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//clean way to create a if-then-else chain with switch without a condition, its a switch true
	fmt.Println("switch with no condition")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("----------------------------------------------")
	fmt.Println("Implementatition of defer to execute a function at the end of the scope")
	{

		defer fmt.Println("world-from de defer section")

		fmt.Println("hello")
	}
	fmt.Println("----------------------------------------------")
	fmt.Println("Stacking defers statements to verify the order of its calls")
	fmt.Println("counting")
	{
		for i := 0; i < 5; i++ {
			defer fmt.Println(i)
		}

		fmt.Println("done")
	}

}

//function implementing if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//function implementing an if but this one has variables defined in the scope that
//can only be used in the if scope until the end of the if
//this function is also implementing the else statement, this one has access to the
//variables defined in the if scope
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
