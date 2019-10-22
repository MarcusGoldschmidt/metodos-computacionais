package main

import (
	"fmt"
	"time"
)

func wait(number int) {
	for i := 0; i < number; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	go wait(5)
	go wait(5)

	fmt.Scanln()
}
