package main

import "fmt"

func main() {
	// This For loop will loop through 3 times and print the numbers 1, 2, and 3
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// This for loop will loop through3 times and print the numbers 7, 8, and 9
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	/* This for loop will loop without the break statement. As a result it will
	only print the word "loop" once
	*/
	for {
		fmt.Println("loop")
		break
	}

	// This for loop will run 3 times and print the numbers 1, 3, and 5
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // Continue moves on to the next iteration of the loop
		}
		fmt.Println(n)
	}
}
