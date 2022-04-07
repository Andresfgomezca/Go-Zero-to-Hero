package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//Range for slice
	for i, v := range pow {
		//range that prints the information of the slice, the index is show after **
		//and the value will be printed after the equals.
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//Range for a map
	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` on arrays and slices provides both the
	// index and value for each entry. Above we didn't
	// need the index, so we ignored it with the
	// blank identifier `_`. Sometimes we actually want
	// the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	//ascii
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
