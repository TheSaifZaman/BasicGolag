package main

import (
	"fmt"
	"time"
)

func task() {
	ticker := time.NewTicker(time.Second * 2)
	for range ticker.C {
		fmt.Println("Tick ")
	}
}

func main() {
	go task()
	time.Sleep(time.Second * 10)
}
