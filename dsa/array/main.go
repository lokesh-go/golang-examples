package main

import "fmt"

func main() {
	// Declare array with size
	var a [5]int

	// Assigning values into array
	for i := 0; i < len(a); i++ {
		a[i] = i * 10
	}

	// Print values
	for i := len(a); i > 0; i-- {
		fmt.Println(a[i-1])
	}
}
