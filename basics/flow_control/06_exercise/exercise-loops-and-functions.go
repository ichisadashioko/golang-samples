package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("[Sqrt] z = %v\n", z)
	}

	return z
}

func main() {
	x := 2.0
	a := Sqrt(x)
	b := math.Sqrt(x)
	diff := b - a
	fmt.Printf("Sqrt(%v) = %v\n", x, a)
	fmt.Printf("math.Sqrt(%v) = %v\n", x, b)
	fmt.Printf("math.Sqrt(%v) - Sqrt(%v) = %v\n", x, x, diff)
}
