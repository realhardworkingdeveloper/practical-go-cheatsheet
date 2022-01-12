package main

import (
	"fmt"
)

/// important
/// command must be: go run test-cutom-module.go custom-module.go
/// important
func main() {
	fmt.Println("access cutom module")
	
	var c int
	
	c = Add(10, 6)
	fmt.Printf("add()=%d\n", c)

	c = Subtract(5, 8)
	fmt.Printf("subtract()=%d\n", c)

	c = Multiply(5, 3)
	fmt.Printf("multiply()=%d\n", c)
}