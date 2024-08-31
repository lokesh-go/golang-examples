package dp

func rob2(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	// Two ways to loot house
	// First way - skip first house and loot others
	// Second way - skip last house and loot others
	// Making two array
	skipLastHouse := make([]int, len(nums)-1)
	skipFirstHouse := make([]int, len(nums)-1)

	// Iterates and fill the array
	for i := 0; i < len(nums)-1; i++ {
		skipLastHouse[i] = nums[i]
		skipFirstHouse[i] = nums[i+1]
	}

	// Get loot for two ways
	lootWithSkipLast := robHouses(skipLastHouse)
	lootWithSkipFirst := robHouses(skipFirstHouse)

	return max(lootWithSkipLast, lootWithSkipFirst)
}

func robHouses(houses []int) int {
	if len(houses) < 2 {
		return houses[0]
	}

	maxLoot := make([]int, len(houses))

	// Max loot from first and second house
	maxLoot[0] = houses[0]
	maxLoot[1] = max(houses[0], houses[1])

	// Max loot from other houses
	for i := 2; i < len(houses); i++ {
		maxLoot[i] = max(maxLoot[i-2]+houses[i], maxLoot[i-1])
	}

	return maxLoot[len(houses)-1]
}
