package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your City: ")

	city, _ := reader.ReadString('\n')
	fmt.Println("You live in " + city)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter your country: ")

	country, _ := reader.ReadString('\n')
	fmt.Println("Your country is " + country)

}
