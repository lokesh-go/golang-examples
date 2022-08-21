package main

import pkg "github.com/lokesh-go/golang-examples/print/pkg"

// Package testing file
func main() {
	// Printes some text but cursor stay on same line
	pkg.Print("Hi Gopher")

	// Prints some text and cursor go to new line
	pkg.PrintWithNewLine("Hello World!")
}
