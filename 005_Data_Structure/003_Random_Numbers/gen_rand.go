package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	// rolls a dice (1 to 6)
	dice := rand.Intn(5)
	fmt.Println("Rolled Dice: ", dice+1)

	// generate random float between 0.0 and 1.0
	randFloat := rand.Float64()
	fmt.Println("Random float between 0.0 and 1.0:", randFloat)

	// generate a random 8-digit ID
	id := ""
	for i := 0; i < 8; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}

	// random number
	randomNum := random(0, 10)
	fmt.Printf("Random number: %d\n", randomNum)

}
