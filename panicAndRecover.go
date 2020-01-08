package main

import (
	"fmt"
	"os"
)

func main() {
	/* Comment out inserted panic
	vals := []int{1, 2, 3}
	// This will cause a panic
	v := vals[10]
	fmt.Println(v)
	*/
	file, err := os.Open("No such file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("file opened")
}
