package dp

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	noOfSteps := make([]int, n)

	// Number of steps to reach at 1 & 2 stairs are
	noOfSteps[0] = 1
	noOfSteps[1] = 2

	// Number of steps to reach at Nth stairs
	for i := 2; i < n; i++ {
		// No of steps to reach at 3 = no of steps to reach at 2 + no of steps to reach at 1
		noOfSteps[i] = noOfSteps[i-1] + noOfSteps[i-2]
	}
	return noOfSteps[n-1]
}
