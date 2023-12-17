package main

import "fmt"

func main() {
	//The simplest counter-based iteration, the basic form is:
	//for initialization statement; conditional statement; modified statement {}
	//The heads of these three sections, which are separated by a semicolon; they do not need to be enclosed in parentheses ().

	for x := 0; x < 4; x++ {
		fmt.Printf("iteration x: %d\n", x)
	}

	// define array
	a := []int{1, 2, 3, 4, 5, 6}

	// loop over array
	for i := 0; i < len(a); i = i + 1 {
		fmt.Println("character :", a[i])
	}
}
