package main

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
