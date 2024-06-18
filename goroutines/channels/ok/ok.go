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

	// Create channel
	stream := make(chan int)

	// Close channel to see the result
	close(stream)

	// Reading from channel with two parameters
	s, ok := <-stream

	// ok variable --- is false, indicating that the value we received is the zero
	// value for int, or 0 and not a value placed on the stream
	fmt.Println(s, ok)
}
