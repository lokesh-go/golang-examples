package main

import (
	"fmt"
	"log"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func new() linkedList {
	return linkedList{}
}

func (l *linkedList) insertAtBeginning(data int) {
	// Create a new node with data
	newNode := &node{data: data}

	// Assign first node to newNode -> next
	newNode.next = l.head

	// Assign new node to head which will be first element
	l.head = newNode
}

func (l *linkedList) insertAtLast(data int) {
	// Create a new node with data
	newNode := &node{data: data}

	// If list is empty
	if l.head == nil {
		l.head = newNode
		return
	}

	// Move pointer to last node
	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	// Assign new node to the last
	lastNode.next = newNode
}

func (l *linkedList) insertAtPosition(data, pos int) {
	// Create node with data
	newNode := &node{data: data}

	// If invalid position
	if pos < 0 {
		log.Fatal("position invalid")
	}

	// If position is beginning
	if pos == 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}

	current := l.head
	for i := 0; i < pos-1 && current != nil; i++ {
		current = current.next
	}

	// Checks
	if current == nil {
		log.Fatal("Position out of bounds")
	}

	newNode.next = current.next
	current.next = newNode
}

func (l *linkedList) delete(data int) {
	// Check for empty
	if l.head == nil {
		return
	}

	// If head node has value
	if l.head.data == data {
		l.head = l.head.next
		return
	}

	currentNode := l.head
	for currentNode.next != nil && currentNode.next.data != data {
		currentNode = currentNode.next
	}

	// If not found
	if currentNode.next == nil {
		fmt.Println("Not Found")
		return
	}

	currentNode.next = currentNode.next.next
}

func (l *linkedList) display() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main() {
	// Get new node
	linkedList := new()

	// Insert value at beginning
	// Time Complexity - O(1)
	linkedList.insertAtBeginning(5)

	// Prints
	// Time Complexity - O(n)
	linkedList.display()

	// Insert value at last
	// Time Complexity - O(n)
	linkedList.insertAtLast(10)

	// Prints
	linkedList.display()

	// Insert value at the position
	// Time Complexity - O(n)
	// Time Complexity - O(1) - at the beginning
	linkedList.insertAtPosition(999, 1)

	// Prints
	linkedList.display()

	// Delete a node
	// Time Complexity - O(n)
	// Time Complexity - O(1) - node deleted is head
	linkedList.delete(10)

	// Prints
	linkedList.display()
}
