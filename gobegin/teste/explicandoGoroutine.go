package main

import (
	"fmt"
	"time"
)

func say(s string) {
	defer fmt.Println("Deferido!")
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}
