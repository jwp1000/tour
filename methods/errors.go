package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",
		float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z, newz := 0.0, 1.0
	//fmt.Println(z)
	for z != newz {
		z = newz
		newz -= (z*z - x) / (2 * z)
		//	fmt.Println(newz)
	}
	return z, nil
}

func doit(x float64) {
	if result, err := Sqrt(x); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sqrt of %v is: %v\n", x, result)
	}
}

func main() {
	doit(4)
	doit(-4)

}
