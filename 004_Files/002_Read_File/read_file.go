package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("hello.txt")

	// can file be opened?
	if err != nil {
		fmt.Print(err)
	}

	// convert bytes to string.
	str := string(b)

	//show file data
	fmt.Println(str)
}
