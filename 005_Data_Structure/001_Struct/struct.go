package main

import "fmt"

type House struct {
	noRoom int
	price  float32
	city   string
}

func main() {
	var newHouse House

	newHouse.noRoom = 4
	newHouse.price = 25000.25
	newHouse.city = "Dhaka"

	fmt.Printf("newHouse.noRoom = %d\n", newHouse.noRoom)
	fmt.Printf("newHouse.price = %f\n", newHouse.price)
	fmt.Printf("newHouse.city = %s\n", newHouse.city)
}
