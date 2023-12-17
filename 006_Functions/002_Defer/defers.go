package main

import (
	"fmt"
	"os"
)

func hello() {
	fmt.Println("Hello")
}

func who() {
	fmt.Println("Go")
}

func main() {

	// The normal execution in a Go function is top to bottom, if you had the function below,
	// it would first call who() (top) and then hello (bottom).
	who()
	hello()

	// If you add the defer statement, it will remember to call it after the function is finished.
	defer who()
	hello()

	// Don't Forget to close an open file.
	f, _ := os.Create("hello.txt")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	fmt.Println(f, "hello world")

	// Predict what this code does: Ha Ha...
	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("World")
}
