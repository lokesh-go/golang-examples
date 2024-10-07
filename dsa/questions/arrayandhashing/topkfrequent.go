package main

/*
	1. Calculate frequency of each num in array
		freq[n]++
	2. Make COUNT-NUM array like:
		- array index is representing the frequency (because it will be in the sorted order)
		- and its values will represent which all numbers have that frequency
	3. Iterates from last (because we have to return top frequent num)
		- append into res
		- if len(res) == k return res
*/

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, v := range nums {
		freq[v]++
	}

	countNum := make([][]int, len(nums)+1)
	for num, count := range freq {
		countNum[count] = append(countNum[count], num)
	}

	res := []int{}
	for i := len(countNum) - 1; i >= 0; i-- {
		for _, ele := range countNum[i] {
			res = append(res, ele)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
