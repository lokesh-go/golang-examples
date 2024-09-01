package main

import (
	"fmt"
	"strconv"
)

func kthLargestNumber(nums []string, k int) string {
	resArray := make([]int, 101)

	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		resArray[n]++
	}

	fmt.Println(resArray)

	for i := len(resArray) - 1; i >= 0; i-- {
		if resArray[i] > 0 {
			k = k - resArray[i]
			if k <= 0 {
				return strconv.Itoa(i)
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(kthLargestNumber([]string{"0", "0"}, 2))
}
