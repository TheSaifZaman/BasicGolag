package main

import (
	"fmt"
	"time"
)

// Select waits on multiple channels. This can be combined with goroutines.
// select is like the switch statement, but for channels.
func f1(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "1"
	}
}

func f2(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "2"
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go f1(c1)
	go f2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
	fmt.Println("Goroutines finished.")
}
