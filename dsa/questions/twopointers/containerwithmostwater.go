package twopointers

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
