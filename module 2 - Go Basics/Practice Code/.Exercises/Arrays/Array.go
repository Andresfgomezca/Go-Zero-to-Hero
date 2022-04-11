package main

import "fmt"

func main() {
	defer fmt.Println("end of the program using defer")
	var a [20]string
	b := "string to fill the array with these chars"
	bChars := []rune(b)
	for i := 0; i < len(a); i++ {
		char := string(bChars[i])
		a[i] = char
	}
	//the array a will be filled with the chars of a string in string type
	for _, v := range a {
		fmt.Println(v)
	}
}
