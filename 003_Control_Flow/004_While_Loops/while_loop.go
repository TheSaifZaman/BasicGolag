package main

import "fmt"

func main() {
	i := 1
	maxNum := 20

	// technically go does not have while, but
	// for can be used while in go.

	for i <= maxNum {
		fmt.Println(i)
		i++
	}
}
