package main

import "fmt"

func main() {
	// Use range to sum the numbers in a slice or array
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// In This for loop we use the range to print out the index number of an array or slice
	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index: ", i)
		}
	}

	// Example of range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cantaloupe"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Range iterating over just the keys of the map
	for k := range kvs {
		fmt.Println("Key:", k)
	}

	// Range on strings over Unicode cod points
	for i, c := range "Go" {
		fmt.Println(i, c)
	}
}
