package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println("Arguments with Program Name:", argsWithProg)
	fmt.Println("Arguments without Program Name:", argsWithoutProg)
	fmt.Println("3rd Argument:", arg)
}
