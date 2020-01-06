/* Fiz Buzz Challenge: Print numbers between 1 and 20 with the additional following conditions:
1) if number is divisible by 3 (e.g., 9) print the word "fizz" instead of printing the number
2) if number is divisible by 5 (e.g., 10) print the word "buzz" instead of printing the number
3) if number is divisible by 3 and 5 (e.g., 15) print the word "fizz buzz" instead of printing the number
*/
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 { // For number 15 with && logical operator
			fmt.Println("fizz buzz")
		} else if i%3 == 0 { // For number 3, 6, 9, 12, and 18
			fmt.Println("fizz")
		} else if i%5 == 0 { // For number 5, 10, and 20
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
