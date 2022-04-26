package main

import (
	"fmt"
	"time"
)

func main() {
	go myFunction()
	fmt.Println("Hello from main goroutine!")
	//this delay allows the subroutine to be finished before the program exits
	time.Sleep(1 * time.Second)
}

func myFunction() {
	fmt.Println("Hello from Goroutine!")
}
