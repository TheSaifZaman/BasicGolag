package main

import "fmt"

func main() {
	world := func() string {
		return "world"
	}
	fmt.Print("hello ", world(), "\n")
}
