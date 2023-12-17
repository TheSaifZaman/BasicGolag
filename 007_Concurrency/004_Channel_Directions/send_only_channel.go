package main

import "fmt"

// If you want a Sand only chan, create it as func f(c chan <- string).
func fn(c chan<- string) {
	c <- "send only channel"
}

func main() {
	c := make(chan string, 1)
	fn(c)
	fmt.Println(<-c)
}
