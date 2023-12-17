package main

import (
	"fmt"
	"time"
)

func f1(c chan string) {
	for {
		time.Sleep(10 * time.Second)
		c <- "10 seconds passed"
	}
}

func main() {
	c1 := make(chan string)
	go f1(c1)

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(11 * time.Second):
		fmt.Println("Timeout!")
	}
}
