package slidingwindow

/*
	1. Take right & left two pointers (left = 0, right = left + 1)
	2. Loop right < len(prices)
		If prices is less than tomorrow then will prefer to buy and star calculating the profit
		- if prices[left] < prices[right]
			Calculate profit = prices[right] - prices[left]
			maxProfit = max(maxProfit, profit)
		else
			Otherwise move forward left with right
			left = right

		right++
*/

func maxProfit(prices []int) int {
	maxProfit := 0
	left := 0
	right := left + 1
	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			maxProfit = max(maxProfit, profit)
		} else {
			left = right
		}
		right++
	}
	return maxProfit
}
