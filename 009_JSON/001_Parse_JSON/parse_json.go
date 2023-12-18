package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// define data structure
	type DayPrice struct {
		Gold     int
		Silver   int
		Platinum int
	}

	// json data
	data := `{
		"gold": 1271,
		"silver": 1284,
		"platinum": 1270
	}`

	var obj DayPrice

	// unmarshall it
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	// can access using struct now
	fmt.Printf("Gold     : $%d.00\n", obj.Gold)
	fmt.Printf("Silver   : $%d.00\n", obj.Silver)
	fmt.Printf("Platinum : $%d.00\n", obj.Platinum)
}
