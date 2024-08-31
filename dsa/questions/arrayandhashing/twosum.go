package main

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for ind, num := range nums {
		remainingNum := target - num
		if _, exists := hashMap[remainingNum]; exists {
			return []int{ind, hashMap[remainingNum]}
		}
		hashMap[num] = ind
	}
	return []int{}
}
