package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMAZON":    1699.8,
		"GOOGLE":    1129.19,
		"MICROSOFT": 98.61, // Must have trailing comma in a multiline array
	}

	// Number of litmes
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MICROSOFT"])

	// Get zero value if not found
	fmt.Println(stocks["TESLA"]) // 0

	// Use two value to see if found
	value, ok := stocks["TESLA"]
	if !ok {
		fmt.Println("TESLA not found")
	} else {
		fmt.Println(value)
	}

	//Set
	stocks["TESLA"] = 322.12
	fmt.Println(stocks)

	// Delete
	delete(stocks, "AMAZON")
	fmt.Println(stocks)

	// Single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
