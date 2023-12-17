package main

import "fmt"

//Channels are a way for goroutines to communicate with each-other. A channel type is defined with the keyword chan.
//Because goroutines run concurrently, they can’t simply pass data from one goroutine to another. Channels are needed.
// How do channels work?
// If you type c <- "message", “message” gets send to the channel.
// Then msg := <- c means receive the message and store it in variable msg.

func f(c chan string) {
	c <- "f() was here."
}

func f2(c chan string) {
	msg := <-c
	fmt.Println("f2:", msg)
}
func main() {
	var c = make(chan string)
	go f(c)
	go f2(c)
	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
