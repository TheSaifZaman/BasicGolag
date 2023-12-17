package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(3 * time.Second)
	go func() {
		<-t1.C
		fmt.Println("Timer expired")
	}()

	_, err := fmt.Scanln()
	if err != nil {
		return
	}

	stop := t1.Stop()
	if stop {
		fmt.Println("Timer stopped")
	}
}
