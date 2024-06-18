package ok

import "fmt"

// Example ...
func Example() {
	// Create channel
	msg := make(chan string)

	// Sending msg into channel in another go routine
	go func() {
		msg <- "Hello World"
	}()

	// Reading from channel with two parameters
	m, ok := <-msg
	fmt.Println(m, ok)
}
