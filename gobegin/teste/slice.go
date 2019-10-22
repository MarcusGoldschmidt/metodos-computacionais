package main

import "fmt"

func main() {
	// Entendendo slice
	list := [5]int{9, 8, 7, 6, 5}
	fmt.Println(list)

	a := list[:2]

	fmt.Println("Valor antigo: ", a)

	// Mudando a referencia
	list[0] = 1
	list[1] = 2

	fmt.Println("Valor novo: ", a)

	// len slice
	fmt.Println(len(a), " ", cap(a))
}
