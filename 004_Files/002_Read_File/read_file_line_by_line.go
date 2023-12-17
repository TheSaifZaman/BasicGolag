package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// read file by line into memory
// all files contents is stores in lines[]
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	// Don't forget to close file after reading it.
	// In go we like to handle this first. So, that we don't forget.
	// And with defer we do it in style. Ha Ha...
	defer func(file *os.File) (string, error) {
		err := file.Close()
		if err != nil {
			return "", err
		}
		return "", nil
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// open file for reading line by line
	lines, err := readLines("GoLang\\BasicGolang\\004_Files\\002_Read_File\\hello.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// print file contents
	for i, line := range lines {
		fmt.Println(i, line)
	}
}
