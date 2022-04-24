package main

import "fmt"

func main() {
	//this will be the map with string keys and int values
	map1 := make(map[string]int)
	//array that will fill the empty map
	strings := [8]string{"first", "second", "third", "fifth", "sixth", "seventh", "eighth"}

	for i, v := range strings {
		map1[v] = i + 1
	}
	fmt.Println("array: ", strings)
	fmt.Println("map: ", map1)
}
