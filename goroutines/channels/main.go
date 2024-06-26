package main

import (
	"github.com/lokesh-go/golang-examples/goroutines/channels/ok"
	"github.com/lokesh-go/golang-examples/goroutines/channels/rangeover"
	"github.com/lokesh-go/golang-examples/goroutines/channels/simple"
)

func main() {
	// Simple channel example
	simple.Example()

	// Reading data from channel with two parameters
	ok.Example()

	// Channel range example
	rangeover.Example()

	// Deadlock example
	//deadlock.Example()
}
