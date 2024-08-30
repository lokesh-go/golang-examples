package main

func lengthOfLongestSubstring(s string) int {
	var longest int
	for ind := range s {
		left := ind
		hashMap := map[byte]struct{}{}
		var length int
		for left < len(s) {
			if _, exists := hashMap[s[left]]; exists {
				break
			}
			hashMap[s[left]] = struct{}{}
			length++
			left++
		}
		longest = max(longest, length)
	}
	return longest
}
