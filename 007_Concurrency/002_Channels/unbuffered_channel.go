package main

import (
	"fmt"
	"time"
)

// In an unbuffered channel, each send operation must be matched by a reception operation,
// making it a synchronous communication mechanism.
func worker(done chan bool) {
	fmt.Println("Worker started")
	time.Sleep(time.Second * 2) // simulate some work
	fmt.Println("Worker finished")

	done <- true // send a signal to the channel that the work is done
}

func main() {
	done := make(chan bool)

	go worker(done) // start a worker goroutine

	<-done // wait for the worker goroutine to finish and receive the signal
	fmt.Println("Main function exiting")
}

//In this example, the worker() function simulates some work by sleeping for 2 seconds.
//Once the work is done, it sends a signal to the done channel by writing true to it.
//The main() function creates the done channel and starts a worker goroutine by calling go worker(done).
//It then waits for the worker to finish by receiving a signal from the done channel using the <-done syntax.
