package simple

import "fmt"

// Example ...
func Example() {
	// Creates hello channel
	helloMsgStream := make(chan string)

	// Create go routine to send the msg
	go func() {
		defer close(helloMsgStream)
		helloMsgStream <- "Hello World"
	}()

	// Prints the msg
	fmt.Println(<-helloMsgStream)
}
