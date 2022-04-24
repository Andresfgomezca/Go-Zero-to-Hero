package main

import (
	"fmt"
	"math"
)

type figure struct {
	ftype string
	//the size of the segments
	side float64
}

//triangle, quadrilateral, pentagon, hexagon, octago
//this functions does not have pointers because it is not required to modify the figure to calculate 
//an area
func area(f figure) float64 {
	return1 := float64(0)
	switch f.ftype {
	case "triangle":
		return1 = f.side * (math.Sin((60*math.Pi)/180) * f.side)
	case "quadrilateral":
		return1 = f.side * f.side
	case "pentagon":
		aphothem := aphothem(f.side, 5)
		return1 = 5 * math.Pow(aphothem, 2) * math.Sqrt(5-(2*math.Sqrt(5)))
		//different ways to calculate the area
		//float64(1.72048)*f.side*f.side
		//(5 / 2) * aphothem * f.side
	case "hexagon":
		return1 = (3 * math.Sqrt(3) * math.Pow(f.side, 2)) / 2
	case "heptagon":
		aphothem := aphothem(f.side, 7)
		return1 = ((7 * f.side * aphothem) / 2)
	case "octagon":
		aphothem := aphothem(f.side, 8)
		return1 = (4 * aphothem * f.side)
	}
	return return1
}

func perimeter(f figure) float64 {
	return1 := float64(0)
	switch f.ftype {
	case "triangle":
		return1 = f.side * 3
	case "quadrilateral":
		return1 = f.side * 4
	case "pentagon":
		return1 = f.side * 5
	case "hexagon":
		return1 = f.side * 6
	case "heptagon":
		return1 = f.side * 7
	case "octagon":
		return1 = f.side * 8
	}
	return return1
}

//this function will calculate the aphothem for the figures
func aphothem(f float64, sides float64) float64 {
	tempVar := float64(((180 / sides) * float64(math.Pi)) / 180)
	return (f / (math.Tan(tempVar) * 2))
}

func main() {

	triangle := figure{ftype: "triangle", side: 5}

	cuad := figure{ftype: "quadrilateral", side: 5}

	penta := figure{ftype: "pentagon", side: 5}

	hex := figure{ftype: "hexagon", side: 5}

	hep := figure{ftype: "heptagon", side: 5}

	octa := figure{ftype: "octagon", side: 5}

	fmt.Println("triangle area: ", area(triangle))
	fmt.Println("triangle perimeter: ", perimeter(triangle))

	fmt.Println("cuad area: ", area(cuad))
	fmt.Println("cuad perimeter: ", perimeter(cuad))

	fmt.Println("penta area: ", area(penta))
	fmt.Println("penta perimeter: ", perimeter(penta))

	fmt.Println("hex area: ", area(hex))
	fmt.Println("hex perimeter: ", perimeter(hex))

	fmt.Println("hep area: ", area(hep))
	fmt.Println("hep perimeter: ", perimeter(hep))

	fmt.Println("octa area: ", area(octa))
	fmt.Println("octa perimeter: ", perimeter(octa))

}
