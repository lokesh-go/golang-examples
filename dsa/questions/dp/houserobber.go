package dp

func rob(nums []int) int {
	// If only one house
	if len(nums) < 2 {
		return nums[0]
	}

	// Calculate max loot at first two house
	maxLoot := make([]int, len(nums))

	// Till first house max loot
	maxLoot[0] = nums[0]

	// Till second house max loot
	maxLoot[1] = max(nums[0], nums[1])

	// Calculate max loot for other houses
	for i := 2; i < len(nums); i++ {
		maxLoot[i] = max(maxLoot[i-2]+nums[i], maxLoot[i-1])
	}
	return maxLoot[len(nums)-1]
}

/*
func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
*/
