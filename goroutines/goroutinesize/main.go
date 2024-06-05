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
	ch := make(chan interface{})
	wg.Add(numGoroutines)
	for i := numGoroutines; i > 0; i-- {
		go emptyGoroutine(&wg, ch)
	}
	close(ch)
	wg.Wait()

	after := memConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

func emptyGoroutine(wg *sync.WaitGroup, ch <-chan interface{}) {
	defer wg.Done()
	<-ch
}

func memConsumed() uint64 {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}
