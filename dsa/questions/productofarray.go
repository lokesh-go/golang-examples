package main

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
