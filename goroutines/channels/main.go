package main

import (
	"github.com/golang-examples/goroutines/channels/deadlock"
	"github.com/golang-examples/goroutines/channels/simple"
)

func main() {
	// Simple channel example
	simple.Example()

	// Deadlock example
	deadlock.Example()
}
