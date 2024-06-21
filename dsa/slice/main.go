package main

/*

Understanding Slices as Reference Types

Slice Header:
	A slice in Go is a three-word data structure:
		- A pointer to the underlying array.
		- The length of the slice.
		- The capacity of the slice (which is the maximum length the slice can reach without allocating more memory).
	When you pass a slice to a function, you are passing a copy of the slice header. This copy still points to the same underlying array as the original slice.

Modifications in Functions:
	- If you modify elements of the slice within a function, those modifications affect the original slice because both the original and copied slice headers refer to the same underlying array.
	- However, if you modify the slice itself (like appending elements), the copied slice header does not affect the original slice header.
	This is because appending may require allocating a new underlying array if the capacity is exceeded, and the copied slice header doesn't have access to the original slice's capacity information.

*/

import (
	"fmt"
	"log"
)

func main() {
	// Initialise slice with some value
	arr := []int{10, 12, 14, 16, 18}

	fmt.Println("Original Arr: ", arr)

	// Modify element
	modifyArray(arr)

	fmt.Println("Modified Arr: ", arr)

	// Insert value at beginning
	// Time Complexity : O(n)
	insertAtBeg := insertAtBeginning(arr, 1)

	fmt.Println("Original Arr (After append): ", arr)

	fmt.Println("Insert At Beg: ", insertAtBeg)

	// Insert value at last
	// Time Complexity : O(1)
	insertAtLast := insertAtLast(arr, 100)

	fmt.Println("Insert At Last: ", insertAtLast)

	// Insert value at middle
	// Time Complexity : O(n)
	insertAtMid := insertAtMiddle(arr, 500, 4)

	fmt.Println("Insert At Mid: ", insertAtMid)
}

func modifyArray(arr []int) {
	// Modify the element
	// Changes will happen in original array (because it pointed to the same underlying array)
	arr[3] = 0
}

func insertAtBeginning(arr []int, val int) []int {
	// Changes will not happen in original array (because appending require allocating a new underlying array)
	arr = append([]int{val}, arr...)
	return arr
}

func insertAtLast(arr []int, val int) []int {
	// Changes will not happen in original array (because appending require allocating a new underlying array)
	arr = append(arr, val)
	return arr
}

func insertAtMiddle(arr []int, val int, index int) []int {
	// Case handling
	if index < 0 || index > len(arr) {
		log.Fatal("Index out of bound")
		return arr
	}

	arr = append(arr[:index], append([]int{val}, arr[index:]...)...)
	return arr
}
