package main

import "fmt"

func main() {
	// A closure is a type of function, that uses variables defined outside the function itself.

	// Ex.1
	x := 2
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	//Ex.2
	sequenceGenerator := makeSequence()
	fmt.Println(sequenceGenerator())
	fmt.Println(sequenceGenerator())
	fmt.Println(sequenceGenerator())
}

func makeSequence() func() int {
	i := 1
	return func() int {
		i += 2
		return i
	}
}
