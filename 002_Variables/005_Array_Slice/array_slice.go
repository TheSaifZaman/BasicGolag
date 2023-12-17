package main

import "fmt"

func main() {
	// define array
	a := [5]int{1, 2, 3, 4, 5}

	//var slice1 []type = make([]type, len)
	//var slice1 []type = make([]type, len, cap)
	//var slice1 []type = arr1[start:end]

	//v := make([]int, 10, 50)
	//This allocates an array with a 50 int value,
	//and creates a slice v with a length of 10 with a capacity of 50,
	//which points to the first 10 elements of the array.

	// create slice from array a[low : high : max] / func make([] T, len, cap)
	t := a[1:3:5]

	// output slice
	fmt.Println(t)
}
