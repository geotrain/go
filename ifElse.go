package main

import "fmt"

func main() {
	/* This If Else Statement divides 7 by 2, if the remainder is equal to 0 the first statement will print
	other wise if it is false the 2nd statement (else) will be printed
	*/
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// This if statement simply prints only if 8 divided by is equal to 0 which it is
	if 8%4 == 0 {
		fmt.Println("8 is dividble by 4")
	}

	// This if, else if, else statement prints the else if line because 9 is not less than 0
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
