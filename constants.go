package main

import "fmt"
import "math"

const s string = "constant" // Declare a constant string

func main() {

	fmt.Println(s)

	const n = 500000000 // Declare a constant integer

	const d = 3e20 / n    // Divide Constant d by 3e20
	fmt.Println(int64(d)) // Print Constant d

	fmt.Println(math.Sin(n)) // Print Trigonometric function Sine and pass it Constant n
}
