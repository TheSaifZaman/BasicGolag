package main

import (
	"fmt"
	"time"
	//"github.com/bmuller/arrow/lib"
)

func main() {
	current := time.Now().UTC()
	fmt.Println("Date: " + current.Format("2006 Jan 02"))
	fmt.Println("Time: " + current.Format("03:04:05"))

	// Strftime formatting
	// fmt.Println("Date: ", arrow.Now().CFormat("%Y-%m-%d"))
	// fmt.Println("Time: ", arrow.Now().CFormat("%H:%M:%S"))
}
