package main

//this code only works in go tour background uncommenting the library and the execution of the test
import (
	"fmt"
	//"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	//this cicle is filling the map with the keys(words) and increasing the
	//value according to the ammount of times that they are repeated
	for _, v := range a {
		m[v]++
	}
	return m
}

func main() {
	fmt.Println("Exercise: Maps")
	//wc.Test(WordCount)

}
