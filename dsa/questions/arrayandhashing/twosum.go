package main

/*
	1. Initialise empty hash map
	2. Range over numbers
	3. Find remaining num (target-num) in hash map
		- If found return (ind of 1st num, ind of 2nd num)
		- else hashMap[num] = ind
*/

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
