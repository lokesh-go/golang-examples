package main

import "sort"

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		// Handling for the different length of string
		// Ideally length of string which greater than other should come later
		l1 := len(nums[i])
		l2 := len(nums[j])

		if l1 != l2 {
			return l1 < l2
		}

		// If have same number of characters length
		// Lexicographical comparison
		return nums[i] < nums[j]
	})

	return nums[len(nums)-k]
}
