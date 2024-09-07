package main

import "fmt"

// Strategy : An interface defining a set of operations (or algorithms)
// Concrete Strategies: Different implementations of the strategy interface

type cache struct {
	storage      map[string]string
	evictionAlgo cacheEvictionAlgo
}

func newCache(e cacheEvictionAlgo) *cache {
	return &cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
	}
}

func (c *cache) addKey(key string, value string) {
	c.storage[key] = value
}

func (c *cache) evict() {
	c.evictionAlgo.evict()
}

type cacheEvictionAlgo interface {
	evict()
}

type fifoEvictionStrategy struct{}

func (f *fifoEvictionStrategy) evict() {
	fmt.Println("fifo eviction algo")
}

type lruEvictionStrategy struct{}

func (l *lruEvictionStrategy) evict() {
	fmt.Println("lru eviction algo")
}

func main() {
	// Init cache
	fifoEvictionStrategy := &fifoEvictionStrategy{}
	cacheWithFifo := newCache(fifoEvictionStrategy)
	cacheWithFifo.evict()

	// Init cache
	lruEvictionStrategy := &lruEvictionStrategy{}
	cacheWithLRU := newCache(lruEvictionStrategy)
	cacheWithLRU.evict()
}
