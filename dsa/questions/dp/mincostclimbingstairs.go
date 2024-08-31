package dp

func minCostClimbingStairs(cost []int) int {
	// Top where we have to reach
	n := len(cost)

	// Taking cost to reach till every step
	minCost := make([]int, n+1)

	// For 1 & 2nd steps we don't have any cost, onward we have to pay cost
	// Because from the ground floor we can easily step into either 1st or 2nd without any cost
	for i := 2; i <= n; i++ {
		// There is two ways to reach at top
		// Either we can take one step to reach at top - cost of one step behind + cost to till one step behind
		// Either we can take two step to reach at top - cost of two steps behind + cost of till two steps behind
		minCost[i] = min(cost[i-1]+minCost[i-1], cost[i-2]+minCost[i-2])
	}
	return minCost[n]
}
