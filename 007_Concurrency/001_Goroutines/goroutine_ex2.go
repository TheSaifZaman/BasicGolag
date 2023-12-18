package main

import (
	"fmt"
	"time"
)

func number() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func letter() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%c ", 'a'+i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go number()
	go letter()
	time.Sleep(1 * time.Second)
	fmt.Println("main")
}
