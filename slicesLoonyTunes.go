package main

import (
	"fmt"
)

func main() {
	//Same Type
	loonyTunes := []string{"Bugs", "Daffy", "Taz"}
	fmt.Printf("loonyTounes = %v (type %T)\n", loonyTunes, loonyTunes)

	//Length
	fmt.Println(len(loonyTunes)) // 3

	fmt.Println("----")
	// 0 indexing
	fmt.Println(loonyTunes[1]) // Daffy

	fmt.Println("----")
	// slices
	fmt.Println(loonyTunes[1:]) // Daffy Taz

	fmt.Println("----")
	// for
	for i := 0; i < len(loonyTunes); i++ {
		fmt.Println(loonyTunes[i])
	}

	fmt.Println("-----")
	// Single value range
	for i := range loonyTunes {
		fmt.Println(i)
	}

	fmt.Println("-----")
	//Double value range
	for i, name := range loonyTunes {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("-----")
	// Double value range, ignore index by using _
	for _, name := range loonyTunes {
		fmt.Println(name)
	}

	fmt.Println("-----")
	// Append
	loonyTunes = append(loonyTunes, "Elmer")
	fmt.Println(loonyTunes) // Bugs Daffy Taz Elmer
}
