package main

import "fmt"

// Declare a function sums() that will take a number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	//Accepting int arguments
	sum(1, 2)
	sum(1, 2, 3)

	// Applying these args into a slice using a variadic function using func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
