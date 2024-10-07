package twopointers

/*
	1. Take two pointers (left = 0, right = len(height)-1)
	2. Loop util left < right
		- containerWid := right - left
		- containerLen := min(height[left], height[right])
		- containerArea := containerWidth * containerLen
		- area = max(area, containerArea)
		- if height of left is less than right
			- left++
		else
			- right--
*/

func maxArea(height []int) int {
	var area int
	left := 0
	right := len(height) - 1
	for left < right {
		containerWidth := right - left
		containerLength := min(height[left], height[right])
		containerArea := containerWidth * containerLength
		area = max(area, containerArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}
