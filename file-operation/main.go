package main

/// should import io/ioutil
import (
	"fmt"
	"io/ioutil"
)

/// main func for test file read/write
func main() {

	/// test write file
	fmt.Println("writing data into a file")
	writeFile("welcome to go")

	/// test read file
	fmt.Println("reading data from a file")
	readFile()
}

/// file write method
func writeFile(message string) {
	bytes := []byte(message)
	ioutil.WriteFile("./test.txt", bytes, 0644)
	fmt.Println("created a file")
}

/// file read method
func readFile() {
	data, _ := ioutil.ReadFile("./test.txt")
	fmt.Println("file content:")
	fmt.Println(string(data))
}