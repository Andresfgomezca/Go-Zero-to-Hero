package main

import (
	"fmt"
	"os"
)

func main() {

	// set environment variable TEST
	os.Setenv("TEST", "test value")

	// returns value of TEST
	fmt.Println("TEST:", os.Getenv("TEST"))

	// Unset environment variable TEST
	os.Unsetenv("TEST")

	// returns empty string and false,
	// because we removed the GEEKS variable
	value, ok := os.LookupEnv("TEST")

	fmt.Println("TEST:", value, " Is present:", ok)

	// list all environment variables and their values
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
