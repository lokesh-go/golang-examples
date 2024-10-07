package main

/*
Find starting point in an array
nums = [100,4,200,1,3,2]
In this example
- Can 100 be a starting point - If 100-1 is not present in the array (To be consecutive we have to find sequential incremental number)
	- Yes it can be because 99 is not in the array
- Can 4 be a starting point - If 4-1 is not present in the array
	- No It can't because 3 is in the array

	1. Take hashMap for numbers which will use to find number in 0(1) time
	2. Iterate over the nums
		- Check if (num-1) exists or not
		- If not then It will be a starting point
			len := 1
			loop util (num + len) exists in the hashMap (consecutive numbers)
				- take max(length)
	3. return length
*/

func longestConsecutive(nums []int) int {
	numMap := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		numMap[n] = struct{}{}
	}

	var longest int
	for _, num := range nums {
		if _, exists := numMap[num-1]; !exists {
			length := 1
			for {
				if _, exists := numMap[num+length]; !exists {
					break
				}
				length++
			}
			longest = max(longest, length)
		}
	}
	return longest
}
