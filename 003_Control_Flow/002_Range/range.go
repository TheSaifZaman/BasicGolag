package main

import "fmt"

func main() {
	//Range iterates over elements. That can be elements of an array, elements of a dictionary or other data structures.
	//When using range, you can name the current element and current index: for i, num := range nums {}.
	//But you are free to ignore the index: for _, num := range nums {}

	nums := []int{1, 2, 3, 4, 5, 6}

	//without index
	for _, num := range nums {
		fmt.Println(num)
	}

	//with index
	for index, element := range nums {
		fmt.Print(index, ") ", element, "\n")
	}
}
