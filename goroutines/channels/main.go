package main

import (
	"github.com/golang-examples/goroutines/channels/ok"
	"github.com/golang-examples/goroutines/channels/simple"
)

func main() {
	// Simple channel example
	simple.Example()

	// Reading data from channel with two parameters
	ok.Example()

	// Deadlock example
	//deadlock.Example()
}
