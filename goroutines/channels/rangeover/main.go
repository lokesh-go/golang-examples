package rangeover

import "fmt"

/*
	Ranging over a channel.
	The `range` keyword --- used in conjunction with the `for` statement --- supports channels as arguments
	Will automatically break the loop when a channel is closed
*/

func Example() {
	// Create channel
	dataStream := make(chan int)

	// Assign data into it
	go insertData(dataStream)

	// Prints channel data using `range` keyword
	for data := range dataStream {
		fmt.Printf("%d ", data)
	}
}

func insertData(dataStream chan<- int) {
	defer close(dataStream)

	for i := 1; i <= 5; i++ {
		dataStream <- i
	}
}
