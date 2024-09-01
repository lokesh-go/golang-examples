package backtracking

func combinationSum(candidates []int, target int) (ans [][]int) {
	sumHelper(candidates, target, 0, 0, &ans, []int{})
	return ans
}

func sumHelper(candidates []int, target int, currentIndex int, currentSum int, answers *[][]int, chooseNumberAtEachLevel []int) {
	// Base case
	if currentSum > target || currentIndex == len(candidates) {
		return
	}

	// If found the sum equals to target
	if currentSum == target {
		*answers = append(*answers, append([]int{}, chooseNumberAtEachLevel...))
		return
	}

	// There is two possiblities
	// Case 1: Either we can not choose a number and move index to ahead
	sumHelper(candidates, target, currentIndex+1, currentSum, answers, chooseNumberAtEachLevel)

	// Case 2: Either we can choose a number
	chooseNumberAtEachLevel = append(chooseNumberAtEachLevel, candidates[currentIndex])
	sumHelper(candidates, target, currentIndex, currentSum+candidates[currentIndex], answers, chooseNumberAtEachLevel)

	// Backtracking : remove last appended value because control is retruning
	chooseNumberAtEachLevel = chooseNumberAtEachLevel[:len(chooseNumberAtEachLevel)-1]
}
