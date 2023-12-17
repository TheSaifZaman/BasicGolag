package main

import "fmt"

//A Golang map is an unordered collection of key-value pairs.
//For every key, there is a unique value. If you have a key, you can look up the value in a map.
//Sometimes it's called an associative array or hash table. An example of a map in Go:
//
//elements := make(map[string]string)
//Its defines as a string to string mapping. In this example we’ll use the periodic elements.
//
//You can map on other variable types too. Below an example of string to int mapping:
//
//alpha := make(map[string]int)
//alpha["A"] = 1

func main() {
	// Expensive Way

	//elements := make(map[string]string)
	//elements["USA"] = "United States of America"
	//elements["UK"] = "United Kingdom"
	//elements["UAE"] = "United Arab Emirates"

	//This will do exactly the same, but is a more elegant notation.
	elements := map[string]string{
		"USA": "United States of America",
		"UK":  "United Kingdom",
		"UAE": "United Arab Emirates",
	}
	fmt.Println(elements)
	fmt.Println(elements["UAE"])

	// Complex One

	website := map[string]map[string]string{
		"Google": {
			"name": "Google",
			"type": "Search",
		},
		"YouTube": {
			"name": "YouTube",
			"type": "video",
		},
	}
	fmt.Println(website)
	fmt.Println(website["Google"]["name"])
	fmt.Println(website["Google"]["type"])
}
