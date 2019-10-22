package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("teste.txt")

	defer f.Close()

	teste, _ := f.Write([]byte("tdsadas"))

	fmt.Println(teste)
}