package main

import (
	"fmt"
	"time"
)

func task(done chan bool) {
	fmt.Print("Task 1 (goroutine) running...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func task2() {
	fmt.Println("Task 2 (goroutine)")
}

func main() {
	done := make(chan bool, 1)
	go task(done)

	fmt.Println("Im in the main thread!")

	if <-done {
		go task2()
		_, err := fmt.Scanln()
		if err != nil {
			return
		}
	}
}
