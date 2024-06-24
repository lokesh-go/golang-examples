package main

import (
	"errors"
	"fmt"
	"log"
)

type stack []int

func new() stack {
	return []int{}
}

// Time Complexity - O(1)
// Push adds an element to the top of the stack
func (s *stack) push(value int) {
	*s = append(*s, value)
}

// Time Complexity - O(1)
// Pop remove and return the top element
func (s *stack) pop() (ele int, err error) {
	if len(*s) == 0  {
		return ele, errors.New("Empty Stack!")
	}
	ele = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ele, nil
}

// Time Complexity - O(1)
// Peek return top element of stack without removing it
func (s *stack) peek() (ele int, err error) {
	if len(*s) == 0 {
		return ele, errors.New("Empty Stack!")
	}
	return (*s)[len(*s)-1], nil
}

// Time Complexity - O(1)
// IsEmpty checks if the stack is empty
func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

// Prints stack elements
func (s *stack) print() {
	for i := len(*s); i > 0; i -- {
		fmt.Printf("\n-----\n| %d |", (*s)[i-1])
	}
}

func main() {
	// Initialise stack
	s := new()

	// Push elements into it
	s.push(5)
	s.push(9)
	s.push(2)
	s.push(7)

	// Prints stack elements
	s.print()

	// Pop element
	popEle, err := s.pop()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("\n\nPop Ele: ", popEle)

	// Prints stack elements
	s.print()

	// Peek element
	peekEle, err := s.peek()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("\n\nPeek Ele: ", peekEle)

	// Prints stack elements
	s.print()
}
