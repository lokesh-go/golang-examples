package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var listLength int
	curr := head
	for curr != nil {
		listLength++
		curr = curr.Next
	}

	detetionNodeIndexFromBeg := listLength - n

	// handling for the head node deletion
	if detetionNodeIndexFromBeg == 0 {
		return head.Next
	}

	ind := 0
	deletionPrevNode := head
	deletionNode := head
	for ind < detetionNodeIndexFromBeg {
		deletionPrevNode = deletionNode
		deletionNode = deletionNode.Next
		ind++
	}

	deletionPrevNode.Next = deletionNode.Next
	return head
}
