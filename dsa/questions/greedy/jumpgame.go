package greedy

func canJump(nums []int) bool {
	// Traverse from the last to check
	// am I able to reach this step from prevous one
	finalDestinationPosition := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		// Do I able to reach to final destination from one step previous
		if i+nums[i] >= finalDestinationPosition {
			// If yes, then now set finaldestination to its previous step
			// and check same again do I able to reach here from previous step
			finalDestinationPosition = i
		}
	}

	// If successfully reach at the first index from the last
	// then we have achieved it
	return finalDestinationPosition == 0
}
