package main

import (
	"fmt"
	"os"
)

func main() {
	// Get string from user
	fmt.Print("Enter text to convert to MP3: \n")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("You entered the following string: ", input)

}
