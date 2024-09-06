package main

// Ensures a class has only one instance and provides a global point of access to that instance.
// Example: Database connection pool, Logger instance.

import (
	"fmt"
	"sync"
)

type single struct{}

// Global instance
var singleInstance *single
var lock sync.Mutex
var once sync.Once

func getInstance() *single {
	if singleInstance == nil {
		// Only lock when checking and creating the instance
		lock.Lock()
		defer lock.Unlock()

		// Double-check if instance is still nil after acquiring the lock
		if singleInstance == nil {
			fmt.Println("Creating a single instance")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}

func getInstanceWithOnce() *single {
	// Thread-safe getInstance using sync.Once
	once.Do(func() {
		fmt.Println("Creating a single instance")
		singleInstance = &single{}
	})

	fmt.Println("Single instance already created")
	return singleInstance
}
