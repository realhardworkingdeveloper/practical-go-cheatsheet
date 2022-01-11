package main

import (
	"fmt"
)

func main() {
	demoPanic()
}

/// panic
func demoPanic() {
	fmt.Println("---demo error handling---")

	/// under defer run completely
	defer func() {
		fmt.Println("do something")
	}()
	/// but panic raise error so stop running
	panic("this is a panic from demonPanic()")

	fmt.Println("This message will never show")
}