package main

import "fmt"

func main() {
	channel := make(chan int, 5)
	channel <- 5
	channel <- 3
	channel <- 9
	close(channel)

	for element := range channel {
		fmt.Println(element)
	}
}
