package main

/// must import "log"
import (
	"fmt"
	"log"
	"os"
)

func main() {
	simpleLogging()
	formattingLogging()
	fileLogging()
}

/// basic log
func simpleLogging() {
	fmt.Println("---simple logging---")
	log.Println("Hello World")
	log.Println("This is a simple error")
}

/// log formatting
func formattingLogging() {
	fmt.Println("---formatting logging---")
	var warning *log.Logger

	warning = log.New(
		os.Stdout,
		"WARNING",
		log.Ldate | log.Ltime | log.Lshortfile)

	warning.Println("This is warning message 1")
	warning.Println("This is warning message 2")
}

/// file log
func fileLogging() {
	fmt.Println("---file logging---")
	file, err := os.OpenFile(
		"./test.log",
		os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Failed to open log file")
		return
	}

	var logFile * log.Logger
	logFile = log.New (
		file,
		"APP: ",
		log.Ldate | log.Ltime | log.Lshortfile)
	
	logFile.Println("This is error message 1")
	logFile.Println("This is error message 2")
	fmt.Println("Done")
}