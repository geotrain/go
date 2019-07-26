package main

import "fmt"

func main() {
	// Example To create an empty map with the make function
	m := make(map[string]int)

	// Set key and value pairs using a name[key] = value syntax
	m["k1"] = 7
	m["k2"] = 13

	// Print a map with the key/value pairs you just created
	fmt.Println("Map:", m)

	// Print out the value for a key with name[key] syntax
	v1 := m["k1"]
	fmt.Println("V1:", v1)

	// Return the length of the map
	fmt.Println("Length:", len(m))

	// Example of the delete to remove key/value pairs from the map
	delete(m, "k2")
	fmt.Println("Map:", m)

	// Second return value when getting a value from a map indicates if the key was present in the map.
	_, prs := m["k2"]
	fmt.Println("Prs:", prs)

	// Declare and initialize a new using jSon syntax
	n := map[string]int{
		"foo": 1,
		"bar": 2}
	fmt.Println("Map:", n)
}
