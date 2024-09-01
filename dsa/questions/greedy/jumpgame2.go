package greedy

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	totalJumps := 0
	reachedHereFromLastJump := 0
	maxCoverageAfterJump := 0

	for ind := 0; ind < len(nums); ind++ {
		maxCoverageAfterJump = max(maxCoverageAfterJump, ind+nums[ind])

		// Window for jump
		// If reached at index which we did from last jump
		// Then perform the next jump
		if reachedHereFromLastJump == ind {
			reachedHereFromLastJump = maxCoverageAfterJump
			totalJumps++

			if maxCoverageAfterJump >= len(nums)-1 {
				// Reached to the destination
				return totalJumps
			}
		}

	}
	return totalJumps
}
