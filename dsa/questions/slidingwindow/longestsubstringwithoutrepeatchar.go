package slidingwindow

/*
	Calculate length in a window util we found duplicates

	loop over string
	assign first index to the left pointer
	from every left pointer to the end of string, we start a window
		- maintain hashMap for this (to check for particular windows char is duplicate or not)
		- If duplicate found then break
		- increase length and left pointer
	- After each window we calculate the max length
	return length
*/

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
