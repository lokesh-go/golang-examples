package slidingwindow

func maxProfit(prices []int) int {
	maxProfit := 0
	left := 0
	right := left + 1
	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			maxProfit = max(maxProfit, profit)
		} else {
			// Buy
			// Because price is low so will buy here
			left = right
		}
		// Sell
		right++
	}
	return maxProfit
}
