package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("file_exists.go"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}

	//Error checking
	if _, err := os.Stat("file-exists2.file"); os.IsNotExist(err) {
		fmt.Printf("File does not exist\n")
	}
}
