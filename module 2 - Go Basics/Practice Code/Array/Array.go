package main

import "fmt"

func main() {
	//simple array of strings
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	//array of int initializated
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println("Length of Primes:", len(primes))
}
