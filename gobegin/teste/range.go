package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for _, i := range pow {
		fmt.Printf("%d\n", i)
	}
}
