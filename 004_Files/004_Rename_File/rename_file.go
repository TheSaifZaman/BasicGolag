package main

import "os"

func main() {
	// source and destination name
	src := "GoLang\\BasicGolang\\004_Files\\003_Write_File\\bd_cities.txt"
	dst := "GoLang\\BasicGolang\\004_Files\\003_Write_File\\cities_of_bd.txt"

	// rename file
	err := os.Rename(src, dst)
	if err != nil {
		return
	}
}
