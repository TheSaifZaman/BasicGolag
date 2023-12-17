package main

import "fmt"

func main() {
	var x int = 5
	fmt.Printf("Address of variable x: %x\n", &x)
	fmt.Printf("Value of variable x: %d\n", x)
}
