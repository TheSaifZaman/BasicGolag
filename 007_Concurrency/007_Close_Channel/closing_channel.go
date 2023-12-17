package main

import "fmt"

func main() {
	c := make(chan int, 5)

	c <- 5

	fmt.Println(<-c)
	defer close(c)
	c <- 3
	fmt.Println(<-c)
}
