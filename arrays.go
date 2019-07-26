package main

import "fmt"

func main() {
	// Array named 'a' with 5 spaces then print out the five empty spaces of 0
	var a [5]int
	fmt.Println("emp", a)

	// This array has the value at an index using the array[index] = value syntact and then get a value
	// with the array[index]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// This returns the length of an array a from the first example
	fmt.Println("len:", len(a))

	// Declaring an array and initializing the value in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	// This is an example of a multi-dimensional array data structure
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
