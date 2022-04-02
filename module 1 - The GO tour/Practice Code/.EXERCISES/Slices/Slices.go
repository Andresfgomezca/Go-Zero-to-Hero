package main

//to make it work, uncomment the library imported and the function that is commented in main

//import (	"golang.org/x/tour/pic")

//Exercise slices will only work in go tour
func Pic(dx, dy int) [][]uint8 {
	//[column][index of the element]type uint8
	slices := make([][]uint8, dy)
	//iterate into the slices
	for i := range slices {
		slices[i] = make([]uint8, dx)
		for j := range slices[i] {
			//operation that interacts with the 3 interesting functions
			slices[i][j] = uint8(((i + j) / 2) - (i * j) - (i ^ j))
		}
	}
	return slices

}

func main() {

	//pic.Show(Pic)

}
