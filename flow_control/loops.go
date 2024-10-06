package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, newz := 0.0, 1.0
	fmt.Println(z)
	for math.Abs(z-newz) > .000001 {
		z = newz
		newz -= (z*z - x) / (2 * z)
		fmt.Println(newz)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(16))
}
