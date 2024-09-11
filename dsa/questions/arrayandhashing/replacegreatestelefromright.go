package main

func replaceElements(arr []int) (res []int) {
	res = make([]int, len(arr))
	res[len(arr)-1] = -1
	maxNum := arr[len(arr)-1]

	for i := len(arr) - 2; i >= 0; i-- {
		res[i] = maxNum
		maxNum = max(maxNum, arr[i])
	}

	return res
}
