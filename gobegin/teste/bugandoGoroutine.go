package main

import (
	"fmt"
)

var teste int = 0

func somaum() {
	teste++
}

func main() {
	// Resultado diferente toda vez
	for i := 0; i < 1000; i++ {
		go somaum()
		fmt.Println(teste)
	}
}