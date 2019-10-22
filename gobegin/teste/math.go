package main

import (
	"fmt"
)

const precision = float64(0.0000000001)

func sqrt(number float64) float64 {
	z := float64(1)
	for !((number-precision) < z*z && z*z < (number+precision)) {
		z -= (z*z - number) / (2 * z)
	}
	return z
}

func teste2(x int)   {

}

var name  = "Global"

func main() {
	b := 5

	fmt.Println(sqrt(81))

	teste2(2)

	fmt.Println(name)
}
