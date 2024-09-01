package main

import (
	"container/heap"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Pop() interface{} {
	popVal := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return popVal
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func findKthLargest(nums []int, k int) int {
	h := &maxHeap{}
	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)
	}

	var kthLargestEle int
	for i := 0; i < k; i++ {
		kthLargestEle = heap.Pop(h).(int)
	}

	return kthLargestEle
}

/*
func main() {
	fmt.Println(kthLargestNumber([]int{0}, 1))
}
*/
