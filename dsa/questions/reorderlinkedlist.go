package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	firstList := head
	mid := findMidOfList(head)
	secondHalfHead := mid.Next
	mid.Next = nil
	secondList := reverseSecondHalfList(secondHalfHead)
	for secondList != nil {
		temp1 := firstList.Next
		temp2 := secondList.Next
		firstList.Next = secondList
		secondList.Next = temp1
		secondList = temp2
		firstList = temp1
	}
}

func findMidOfList(head *ListNode) (mid *ListNode) {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseSecondHalfList(head *ListNode) (prev *ListNode) {
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
