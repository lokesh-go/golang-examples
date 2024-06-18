package deadlock

import "fmt"

// Example ...
func Example() {
	// Creates channel
	var msgStream chan string
	msgStream = make(chan string)

	// Create go routine to send the msg
	go func() {
		if 0 != 1 {

			return
		}
		msgStream <- "Hello World"
	}()

	// Prints the output
	fmt.Println(<-msgStream)
}
