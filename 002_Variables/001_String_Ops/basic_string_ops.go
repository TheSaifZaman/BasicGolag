package main

import (
	"fmt"
	"strings"
)

func main() {
	//String variable
	var str1 = "This is a string variable!"
	fmt.Println(str1)

	str1 = "This is another string variable!"
	fmt.Println(str1)

	fmt.Printf("Hello guys, This is string variable 1. \n And this is another 2.")

	//String index
	s := "I am GoLang."
	for k, v := range s {
		fmt.Printf("k: %d, v: %c == %d\n", k, v, v)
	}

	//Direct-use operators
	str := "Beginning of the string" + "End of the string"
	fmt.Println(str)

	dst := "Hell" + "o,"
	dst += " World!"
	fmt.Println(dst)

	//String join
	s = strings.Join([]string{"Hello", "World!"}, ",")
	fmt.Println(s)

	//Print variables
	fmt.Printf("%d:%s", 2023, "year")
}
