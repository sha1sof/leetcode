package task5

func maxProfit(prices []int) int {
	result := 0

	if len(prices) == 0 {
		return 0
	}

	noStock := 0
	profit := -prices[0]

	for _, price := range prices[1:] {
		tempNoStockProfit := result

		result = max(result, profit+price)

		profit = max(profit, noStock-price)

		noStock = tempNoStockProfit
	}

	return result
}
