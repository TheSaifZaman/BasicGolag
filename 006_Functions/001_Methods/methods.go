package main

import "fmt"

func main() {
	hello("go")
	hello("")

	var a float64 = 3
	var b float64 = 9
	var ret = multiply(a, b)
	fmt.Printf("Value is : %.2f", ret)
}

func hello(s string) {
	fmt.Printf("Hello %s", s)
}

func multiply(num1, num2 float64) float64 {
	var result float64
	result = num1 * num2
	return result
}
