package main

import (
	"fmt"
)

func main() {
	v := "Fora"
	fmt.Println(v)
	{
		v := "Dentro"
		fmt.Println(v)
		{
			fmt.Println(v)
		}
	}
	{
		fmt.Println(v)
	}
	fmt.Println(v)
	// Fora
	// Dentro
	// Dentro
	// Fora
	// Fora
}
