package binarysearch

import "math"

func findMin(nums []int) int {
	minNum := math.MaxInt
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[low] <= nums[mid] {
			minNum = min(minNum, nums[low])
			low = mid + 1
		} else {
			minNum = min(minNum, nums[mid])
			high = mid - 1
		}
	}
	return minNum
}
