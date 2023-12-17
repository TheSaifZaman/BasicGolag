package main

import "os"

// Write file in go. You don't need to set any flags.
// This will overwrite the file if it already exists
func main() {
	file, err := os.Create("GoLang\\BasicGolang\\004_Files\\003_Write_File\\cities_of_bd.txt")
	if err != nil {
		return
	}

	// Don't forget to close file after reading it.
	// In go we like to handle this first. So, that we don't forget.
	// And with defer we do it in style. Ha Ha...
	defer func(file *os.File) (string, error) {
		err := file.Close()
		if err != nil {
			return "", err
		}
		return "", nil
	}(file)

	_, err = file.WriteString("These are the cities of Bangladesh: \nDhaka \nRajshahi \nCox'sBazar")
	if err != nil {
		return
	}
}
