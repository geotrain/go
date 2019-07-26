package main

import "fmt"

//Declare functions called Plus
func plus(a int, b int) int { // pass in Two integers with variable names a, b
	// Return results of a + b
	return a + b
}

// Declare function plusPlus() passing in variables a, b, c
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Now call a function with name(aruments) inside main function

func main() {

	result := plus(1, 2)
	fmt.Println("1 + 2 =", result)

	result = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", result)
}
