package main

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
