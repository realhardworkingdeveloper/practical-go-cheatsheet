package main

import (
	"fmt"
)

func main() {
	fmt.Println("simple channel")

	/// define a channel
	c := make(chan int)

	/// run a function in background
	go func() {
		fmt.Println("goroutine process")

		c <- 10 /// write data to a channel
	} ()

	val := <- c
	fmt.Printf("value: %d\n", val)
}