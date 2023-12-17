package main

import "fmt"

// var var_name *var-type
func main() {
	// variable
	var x int = 5

	// create pointer
	var ipointer *int

	// store the address of x in pointer variable
	ipointer = &x

	// display info
	fmt.Printf("Memory address of variable x: %x\n", &x)
	fmt.Printf("Memory address stored in ipointer variable: %x\n", ipointer)
	fmt.Printf("Contents of *ipointer variable: %d\n", *ipointer)
}
