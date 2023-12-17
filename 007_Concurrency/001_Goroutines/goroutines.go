package main

import "fmt"

// A goroutine is a function that can run concurrently. You can see it as a lightweight thread.
// The idea stems from concurrency: working on more than one task simultaneously.

// To invoke a Go routine, write go before the function call.

// If you have a function f(string), call it as go f(string) to invoke it as goroutine.
// The function will then run asynchronously.
func f(msg string) {
	fmt.Println(msg)
}

func main() {
	go f("go routine")
	f("function")
	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
