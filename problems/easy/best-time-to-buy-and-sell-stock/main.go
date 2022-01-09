package easy_best_time_to_buy_and_sell_stock

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}
