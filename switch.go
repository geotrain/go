package main

import "fmt"
import "time"

func main() {
	// This is a basic switch statement that will print case 2
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// This switch statement will use the time function to determine if it prints out weekend or weekday.
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("It's the weekend.")
	case time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's a weekday.")
	}

	// This Switch statement for determining what time it is.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before 12 noon.")
	default:
		fmt.Println("It's after 12 noon.")
	}

	// This switch statement compares types to see if it is boolean (true / false) or not, then tells if it is an
	// integer or not and finally deterines it doesn't know that it's a type string
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
