package main

import "fmt"

// In contrast to unbuffered channels, buffered channels allow you to send a specified
// number of values without needing to wait for a reception operation.
func producer(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i // send the value to the channel
	}
	close(c) // close the channel after all values have been sent
}

func main() {
	c := make(chan int, 2) // create a buffered channel with capacity 2
	go producer(c)

	for i := range c { // receive the values from the channel
		fmt.Println(i)
	}

	// Buffered Channel
	var ch = make(chan string, 3)

	ch <- "hello"
	ch <- "world"
	ch <- "go"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//In this example, the producer() function sends 5 values to the channel c.
//Since c is a buffered channel with capacity 2, the first 2 values are sent immediately,
//and the remaining values are buffered until they can be received by the main() function’s loop.
//Finally, the channel is closed to signal that all values have been sent.
//The main() function receives the values from the channel using the range syntax,
//which automatically stops when the channel is closed.
