package main

import "fmt"

func main() {
	slice := make([]int, 0)

	slice = append(slice, 1, 2, 3, 4)

	// Passandoa  referencia
	refSlice := slice
	refSlice[0] = 100

	fmt.Println("Slice: ", slice)
	fmt.Println("RefSlice: ", refSlice)

	// Copiando o conteudo do ponteiro
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)

	newSlice[0] = 500

	fmt.Println("\nSlice: ", slice)
	fmt.Println("NewSlice", newSlice)

	slice2D := make([][]int, 3)

	for i := 0; i < 3; i++ {
		slice2D[i] = make([]int, i+1)

		for j := 0; j < i+1; j++ {
			slice2D[i][j] = j + i
		}
	}
	fmt.Println("\n", slice2D)
}
