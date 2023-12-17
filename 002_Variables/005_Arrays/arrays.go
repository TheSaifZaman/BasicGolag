package main

import "fmt"

func main() {
	var arr1 = new([5]int)
	fmt.Println(arr1[0])

	//Several assignment modes:
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrKeyValue = [2]string{0: "Chris", 1: "Ron"}

	fmt.Println(arrAge)
	fmt.Println(arrLazy)
	fmt.Println(arrKeyValue)
}
