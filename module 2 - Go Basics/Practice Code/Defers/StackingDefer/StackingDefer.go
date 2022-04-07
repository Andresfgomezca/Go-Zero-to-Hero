package main

import "fmt"

func main() {
	fmt.Println("counting")
	//The program is stacking the defer requests, once it prints done it will start executing the
	//requests in the stack from the last to the first request.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
