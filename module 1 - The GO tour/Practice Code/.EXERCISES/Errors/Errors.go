package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not process the sqrt for the negative number:%v", float64(e))
}

//the function now returns the error value if the sqrt is a complex number.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		if math.Abs(z-(z-(z*z-x)/(z*2))) < 0.00000000000001 {
			return z, nil
		} else {
			z = z - (z*z-x)/(z*2)
		}
	}
}

func main() {
	//error test that returns nil
	fmt.Println(Sqrt(2))
	//error test that returns 0 and the error message.
	fmt.Println(Sqrt(-2))
}
