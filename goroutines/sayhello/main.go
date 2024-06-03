package main

import (
	"fmt"
	"sync"
)

/*
Go follows a model of concurrency called the FORK-JOIN model.
Fork - Split off a child branch of execution to be run concurrently with its parent.
Join - At some point in the future, these concurrent branches of execution will join back together.
Where the child rejoins the parent is called a JOIN POINT.
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// Creates child branch
	go sayHello(&wg)

	// This is JOIN POINT
	wg.Wait()
}

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}
