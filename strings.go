//Go strings example
package main

import (
	"fmt"
)

func main() {
	book := "The color of magic"
	fmt.Println(book)

	fmt.Println(len(book)) // len() determines how many lines are in a string

	fmt.Printf("book[0] = %v, (type %T)\n", book[0], book[0])
	// uint8 is a byte

	//String in go are immutable (cannot change them once declared)

	//Slice (start, end) 0 based, 1/2 empty range - this means printing starts at index 4 but not print index 11
	fmt.Println(book[4:11])

	//Slice (no end) - printing starts at index 4 through the end
	fmt.Println(book[4:])

	//Slicke no start - start at index 0 and go through but not print index 4
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	//Unicode
	fmt.Println("It was 1/2 price!")

	// Multi line
	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`
	fmt.Println(poem)
}
