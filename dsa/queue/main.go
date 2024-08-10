package main

import (
	"errors"
	"fmt"
)

type queue struct {
	data chan int
}

func new(size int) *queue {
	// Creates channel with declared size
	ch := make(chan int, size)

	// Returns
	return &queue{data: ch}
}

// Push
// Enqueue : O(1)
func (q *queue) push(v int) (err error) {
	select {
	case q.data <- v:
		return nil
	default:
		return errors.New("queue is full")
	}
}

// Pop
// Dequeue : O(1)
func (q *queue) pop() int {
	return <-q.data
}

// Size
// Returns the number of elements in the queue
func (q *queue) size() int {
	return len(q.data)
}

// IsEmpty
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

// Clear
func (q *queue) clear() {
	for len(q.data) > 0 {
		<-q.data
	}
}

func main() {
	// Initialise queue
	q := new(7)

	// Push into the queue
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	q.push(5)

	// Pop
	fmt.Println("Pop element: ", q.pop())
	fmt.Println("Pop element: ", q.pop())
	fmt.Println("Pop element: ", q.pop())

	// size
	fmt.Println("Size: ", q.size())

	// Checks for empty
	fmt.Println("IsEmpty: ", q.isEmpty())
}
