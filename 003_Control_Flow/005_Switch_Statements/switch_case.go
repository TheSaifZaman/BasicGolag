package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter integer: ")

	intNum, _ := reader.ReadString('\n')

	switch a, _ := strconv.Atoi(intNum); {
	case a == 1:
		fmt.Println("The integer was == 1")
		fallthrough
	case a == 2:
		fmt.Println("The integer was == 2")
	case a == 3:
		fmt.Println("The integer was == 3")
		fallthrough
	case a == 4:
		fmt.Println("The integer was == 4")
	case a == 5:
		fmt.Println("The integer was == 5")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
