package pkg

import "fmt"

// Prints value but cursor stay on same line
func Print(a interface{}) {
	fmt.Print(a)
}

// Prints value and cursor go to the new line
func PrintWithNewLine(a interface{}) {
	fmt.Println(a)
}
