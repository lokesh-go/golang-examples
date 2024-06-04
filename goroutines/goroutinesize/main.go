package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	numGoroutines = 1e6
)

func main() {
	before := memConsumed()

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for i := numGoroutines; i > 0; i-- {
		go emptyGoroutine(&wg)
	}
	wg.Wait()

	after := memConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

func emptyGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()
}

func memConsumed() uint64 {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}
