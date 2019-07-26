package main

import "fmt"

// Declare function valu() with integer inputs
func vals() (int, int) {
	return 3, 7 // Return results
}

func main() {

	// Here will use 2 different values from the call with multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Using a blank identifier _. we will subet the returned values
	_, c := vals()
	fmt.Println(c)
}
