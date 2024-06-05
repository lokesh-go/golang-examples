package main

import (
	"fmt"
	"sync"
)

/*
Mutex: Allows only one goroutine to access the critical section at a time, either for reading or writing.
RWMutex: Allows multiple goroutines to read simultaneously but only one goroutine to write.
This improves performance when reads are more frequent than writes.
*/

var (
	count   int
	rwMutex sync.RWMutex
)

func main() {
	var wg sync.WaitGroup

	// Reader
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rwMutex.RLock()
			fmt.Println("Read Counter: ", count)
			rwMutex.RUnlock()
		}()
	}

	// Increment
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rwMutex.Lock()
			count++
			rwMutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter: ", count)
}
