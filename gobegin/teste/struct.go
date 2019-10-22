package main

import (
	"fmt"
	"math"
)

type Cartesian struct {
	X int
	Y int
}

func (r Cartesian) Vector() float64 {
	return math.Sqrt(math.Pow(float64(r.Y), 2) + math.Pow(float64(r.X), 2))
}

func main() {

	coordinates := Cartesian{}
	coordinates.Y = 9

	p := &coordinates
	p.X = 5

	fmt.Println(coordinates)

	fmt.Println(p.Vector())
}
