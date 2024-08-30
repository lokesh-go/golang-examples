package main

func characterReplacement(s string, k int) int {
	var res int
	hashMap := make(map[int]int, 26)
	var maxFreqCount int
	var left, right int
	for right < len(s) {
		hashMap[int(s[right]-'A')]++
		maxFreqCount = max(maxFreqCount, hashMap[int(s[right]-'A')])
		widowSize := right - left + 1
		replacementNeeds := widowSize - maxFreqCount
		if replacementNeeds <= k {
			res = max(res, widowSize)
		} else {
			hashMap[int(s[left]-'A')]--
			left++
		}
		right++
	}
	return res
}
