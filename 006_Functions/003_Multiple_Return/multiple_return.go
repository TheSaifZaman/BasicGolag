package main

import "fmt"

func values() (int, string, bool) {
	return 2, "hi", true
}

func main() {
	x, y, z := values()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
