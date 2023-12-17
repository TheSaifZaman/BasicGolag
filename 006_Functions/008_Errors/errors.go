package main

import (
	"errors"
	"fmt"
)

func do() (int, error) {
	return -1, errors.New("something wrong")
}

func main() {
	r, e := do()
	if r == -1 {
		fmt.Println(e)
	} else {
		fmt.Println("Everything went fine.")
	}
}
