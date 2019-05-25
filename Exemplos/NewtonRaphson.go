package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return (x * x * x) - (9 * x) + 3
}

// Derivada
func fLinha(x float64) float64 {
	return (3 * x * x) - 9
}

func CalculaX(x float64) float64 {
	return x - (f(x) / fLinha(x))
}

func main() {

	var x = float64(1)
	err := float64(1) / 100

	for math.Abs(f(x)) > err {
		fmt.Println(x, f(x))
		x = CalculaX(x)
	}
}
