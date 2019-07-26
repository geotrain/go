package main

import "fmt"

func main() {
	// Create an empty string slice with the length of 3
	s := make([]string, 3)
	fmt.Println("Empty:", s)

	// Set and get using arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set: ", s)
	fmt.Println("Get: ", s[2])

	// Return the length of the slice s
	fmt.Println("length: ", len(s))

	// Appending values to a slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Print("Append:", s, "\n")

	// Copying the contents of one slice into an empty slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy: ", c)

	// Example of slice[low:high] to print out s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("Slice 1:", l)

	// Example of Slice up to but excluses s[5]
	l = s[:5]
	fmt.Println("Slice 2:", l)

	// Example of slices up from and includes s[2]
	l = s[2:]
	fmt.Println("Slice 3:", l)

	// Declare and initialize a varialbe for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("Declare:", t)

	// Example of slices as a multi-dimensional data structure
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two Dimensional: ", twoD)
}
