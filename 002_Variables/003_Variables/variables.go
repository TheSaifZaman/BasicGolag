package main

import (
	"fmt"
	"time"
)

func main() {
	//One Different at a time

	//var x int
	//var y bool
	//var z string

	//Multiple Different at a time

	//var (
	//	an int
	//	b bool
	//	s string
	//)

	//	When a variable is declared by a var, the system automatically gives it a zero value of that type.
	//	int i //= 0
	//	float f //= 0.0
	//	bool b //= False
	//	string s //= " "
	//	pointer p // = nil

	//Multiple variables can be assigned on the same line, also known as parallel assignments. For example:
	//
	//a, b, c = 5, 7, "abc"

	// simple declaration:
	//a, b, c := 5, 7, "abc"

	//Scoped Variable
	x := 1
	fmt.Println(x) // prints 1
	{
		//Inside a block, we can redefine the same variable with different value and use them
		//& when coming out of scope it take backs its original value.
		fmt.Println(x) // prints 1
		x := 2
		fmt.Println(x) // prints 2
	}
	fmt.Println(x)

	//String as Variable
	h := "Hello"
	w := "World"

	fmt.Printf("%s %s\n", h, w)

	//Numeric Variables

	//Integers
	i := 1
	fmt.Printf("i is %d\n", i)

	//Integers use the decimal system (base 10), but you can output integers in other ways like binary (base 2),
	//octal (base 8) or hexadecimal (base 16):

	//%b  base 2
	//%c  the character represented by the corresponding Unicode code point
	//%d  base 10
	//%o  base 8
	//%O  base 8 with 0o prefix
	//%q  a single-quoted character literal safely escaped with Go syntax.
	//%x  base 16, with lower-case letters for a-f
	//%X  base 16, with upper-case letters for A-F
	//%U  Unicode format: U+1234; same as "U+%04X"

	//Float
	apple := 3.0
	bread := 2.0
	price := apple + bread

	fmt.Printf("")
	fmt.Printf("Price:    %f\n", price)
	vat := price * 0.15
	fmt.Printf("VAT:      %f\n", vat)
	total := vat + price
	fmt.Printf("Total:    %f\n", total)
	fmt.Printf("")
	//Note: In the above output there are many numbers after the dot. You can limit this by adding a number.
	//For 2 digits %.2f, for 3 digits %.3f etc.

	//Variable swapping made easy in Golang
	g := 1
	f := 2
	fmt.Println(g)
	fmt.Println(f)
	g, f = f, g
	fmt.Println(g)
	fmt.Println(f)

	//nil value
	//The nil designator is used to represent a “zero value” of interface, function, maps, slices, and channels.
	//If you do not specify the type of variable, the compiler will not be able to compile your code because it cannot guess a specific type.

	var num int
	_ = num

	//adding an element in a slice of nil is no problem, but doing the same thing to one map will generate a runtime panic:

	var m map[string]int

	m["one"] = 1 //error

	//string is not nil
	//This is a place that requires attention for developers who often use nil assignment string variables.
	//var str string = " "
	//" " is the zero value of the string and the writing below is the same as the above:
	//var str string

	calculateBirthYear()
}

func calculateBirthYear() {
	t := time.Now()
	fmt.Printf("You were born in %d.", t.Year()-27+1)
}
