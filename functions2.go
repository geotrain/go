// Basic function definition
package main

import (
	"fmt"
)

// Add A to B
func add(a int, b int) int {
	return a + b
}

//div mode returns quotient and reminder
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)
}
