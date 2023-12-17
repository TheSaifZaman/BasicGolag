package main

import "fmt"

// You can define a channel as func f(c <- chan string) to have a receive-only chan.
func f(c <-chan string) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan string, 1)
	c <- "hello"
	f(c)
}
