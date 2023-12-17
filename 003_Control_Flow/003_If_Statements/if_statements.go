package main

import "fmt"

func main() {
	var x = 3

	if x > 2 {
		fmt.Printf("x is greater than 2")
	} else if x == 2 {
		fmt.Printf("condition is true (x = 2)")
	} else {
		fmt.Printf("condition is false (x > 2)")
	}
}
