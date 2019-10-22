package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return (x * x * x) - (9 * x) + 3
}

func CalculaX(a float64, b float64) float64 {
	return (a + b) / 2
}

func CalculaXposisaoFalse(a float64, b float64) float64 {
	return ((a * f(b)) - (b * f(a))) / (f(b) - f(a))
}

func main() {
	a := float64(1)

	b := float64(-1)

	x := CalculaX(a, b)

	err := float64(1) / 100

	if f(a)*f(b) < 0 {
		for math.Abs(f(x)) > err {
			x = CalculaX(a, b)

			if f(x) < 0 && f(a) < 0 {
				a = x
			} else {
				b = x
			}

		}
		fmt.Println(x, f(x))
	} else {
		fmt.Println("Intervalo não contem zero")
	}
}
