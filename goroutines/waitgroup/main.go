package main

import (
	"fmt"
	"sync"
)

/*
Waitgroup is a great way to wait for a set of concurrent operations to complete.
Waitgroup like a concurrent safe counter: calls to Add increment the counter by the integer passed in
and calls to Done decrement the counter by one.
Calls to Wait block until the counter is zero
*/

func main() {
	var wg sync.WaitGroup

	// Here we call Add with an argument of 1 to indicate that one goroutine is beginning
	wg.Add(1)
	go func() {
		// Done ensure that before we exit the goroutine, we indicate to the WaitGroup that we have exited.
		defer wg.Done()
		fmt.Println("1st goroutine")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine")
	}()

	// Call Wait, which will block the main goroutine until all goroutines have indicated they have exited.
	wg.Wait()

	fmt.Println("All goroutines complete")
}

/*
Noticed that the calls to Add are outside the goroutines they are helping to track.
If we didn't do this, we would have introduced a race condition.
*/
