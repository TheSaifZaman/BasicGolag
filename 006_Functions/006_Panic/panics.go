package main

import (
	"os"
)

func main() {
	_, err := os.Create("example.txt")
	if err != nil {
		panic("Cannot create file")
	}
}
