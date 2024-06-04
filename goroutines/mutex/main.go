package main

import (
	"fmt"
	"sync"
)

/*
Mutex stands for "mutual exclusion" and is a way to guard critical sections
of a program.
A critical section is an area of a program that requires exclusive access
to a shared resource.
A Mutex provides a concurrent safe way to express exclusive access to these
shared resources.
A Mutex shares memory by creating a convention developers must follow
to synchronize access to the memory.
Developer is responsible for coordinating access to this memory by
guarding access to it with a Mutex.
*/

var (
	count int
	mu    sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	// Increase
	for i := 5; i > 0; i-- {
		wg.Add(1)
		go increase(&wg)
	}

	// Decrease
	for i := 5; i > 0; i-- {
		wg.Add(1)
		go decrement(&wg)
	}

	wg.Wait()
	fmt.Println("\nComplete")
}

func increase(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	count++
	fmt.Println("Increase: ", count)
	mu.Unlock()
}

func decrement(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	count--
	fmt.Println("Decrease: ", count)
	mu.Unlock()
}
