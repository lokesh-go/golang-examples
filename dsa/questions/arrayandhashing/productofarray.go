package main

/*
	The trickier part is when array contains 0
	1. Counts the zero & calculate the product of nums in the array (if num is 0 don't multiply just continue)
	2. if zeroCount > 1
		return array with all 0 nums -> make([]int, len(nums))
	3. Forms res
		- If zeroCount == 1
			- Case 1:
					num is 0 -> append(res, product)
			- Case 2:
					rest of element will be 0
		else
			- res = append(res, product / num)
*/

func productExceptSelf(nums []int) []int {
	product := int(1)
	zeroCount := 0
	for _, v := range nums {
		if v == 0 {
			zeroCount++
			continue
		}
		product = product * v
	}

	if zeroCount > 1 {
		return make([]int, len(nums))
	}

	res := []int{}
	for _, v := range nums {
		if zeroCount == 1 {
			if v == 0 {
				res = append(res, product)
			} else {
				res = append(res, 0)
			}
		} else {
			res = append(res, product/v)
		}
	}
	return res
}
