package main

import (
	"fmt"
)

func main() {
	demoRecover()
}

/// recover
func demoRecover() {
	fmt.Println("----demo error handling---")
	c := 0
	defer func() {
		a := 10
		b := 0
		
		c = a/b

		/// catch error using recover
		if error := recover(); error != nil {
			fmt.Println("Recovering...", error)
			fmt.Println(error)
		}
	} ()
	fmt.Printf("result = %d\n", c)
	fmt.Println("Done")
}