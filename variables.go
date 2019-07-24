package main

import "fmt"

func main() {

	var a = "initial" // Declare string variable named 'a'
	fmt.Println(a)    // Print String Variable

	var b, c int = 1, 2 // Declare two integer variables
	fmt.Println(b, c)   // Print Integer Variables

	var d = true   // Declare boolean variable to true
	fmt.Println(d) // Print boolean variable

	var e int      // Declare integer variable but not setting a value
	fmt.Println(e) // Prints 0 as the value

	f := "apple"   // Another way to declare a string variable
	fmt.Println(f) // Print string variable
}
