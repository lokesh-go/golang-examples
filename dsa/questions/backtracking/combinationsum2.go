package backtracking

import "sort"

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	combinationSum2Helper(candidates, target, 0, &ans, []int{})
	return ans
}

func combinationSum2Helper(candidates []int, target int, currentIndex int, answers *[][]int, chooseNumberAtEachLevel []int) {
	// Found the combinations
	if target == 0 {
		*answers = append(*answers, append([]int{}, chooseNumberAtEachLevel...))
		return
	}

	// Indexes exceeded
	if target < 0 || currentIndex >= len(candidates) {
		return
	}

	for i := currentIndex; i < len(candidates); i++ {
		// Skip duplicates in the same recursive level
		if i > currentIndex && candidates[i] == candidates[i-1] {
			continue
		}

		chooseNumberAtEachLevel = append(chooseNumberAtEachLevel, candidates[i])
		combinationSum2Helper(candidates, target-candidates[i], i+1, answers, chooseNumberAtEachLevel)

		// Backtracking : remove last appended value because control is retruning
		chooseNumberAtEachLevel = chooseNumberAtEachLevel[:len(chooseNumberAtEachLevel)-1]
	}
}
