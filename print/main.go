package main

import printPkg "golang-examples/print/pkg"

// Package testing file
func main() {
	// Printes some text but cursor stay on same line
	printPkg.Print("Hi Gopher")

	// Prints some text and cursor go to new line
	printPkg.PrintWithNewLine("Hello World!")
}
