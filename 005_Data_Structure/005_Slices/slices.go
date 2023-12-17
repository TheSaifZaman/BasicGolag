package main

import "fmt"

func main() {
	/* create a set/list of numbers */
	mySet := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	/* take slice */
	n := mySet[0:4]
	fmt.Println(n)

	/* define a string */
	myString := "Go programming"

	/* take slice */
	s := myString[0:2]
	fmt.Println(s)
}
