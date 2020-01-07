//Find maximum value in an array of integers
package main

import (
	"fmt"
)

func main() {
	numbers := []int{16, 89, 4, 42, 23, 15}

	// Code goes here
	max := numbers[0] // Initialize max with the first value
	// [1:] jskips the first value
	for _, value := range numbers[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)
}
