package main

import (
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) (res [][]int) {
	hashMap := make(map[string]struct{}, len(nums))
	sort.Ints(nums)
	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := num + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				temp := []int{num, nums[j], nums[k]}
				key := getKey(temp)
				if _, exists := hashMap[key]; !exists {
					hashMap[key] = struct{}{}
					res = append(res, temp)
				}
				j++
				k--
				for j < k {
					if nums[j] == nums[j-1] {
						j++
					} else {
						break
					}
				}
				for j < k {
					if nums[k] == nums[k+1] {
						k--
					} else {
						break
					}
				}
			}
		}
	}
	return res
}

func getKey(temp []int) string {
	var strArr []string
	for _, n := range temp {
		strArr = append(strArr, strconv.Itoa(n))
	}
	return strings.Join(strArr, ",")
}
