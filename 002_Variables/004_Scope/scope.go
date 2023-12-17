package main

import "fmt"

// global scope
var x = 7

func example() {
	fmt.Println(x)
}

func main() {
	fmt.Println(x)
	example()
	fmt.Println(x)
	{
		//local scope
		x = 8
		fmt.Println(x)
	}
	fmt.Println(x)
}
