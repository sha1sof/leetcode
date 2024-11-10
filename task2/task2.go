package task2

func buyChoco(prices []int, money int) int {
	if len(prices) < 2 {
		return money
	}
	var min1 int
	var min2 int
	if prices[0] < prices[1] {
		min1 = prices[0]
		min2 = prices[1]
	} else {
		min1 = prices[1]
		min2 = prices[0]
	}
	for i := 2; i < len(prices); i++ {
		if prices[i] < min2 {
			if prices[i] < min1 {
				min2 = min1
				min1 = prices[i]
			} else {
				min2 = prices[i]
			}
		}
	}

	if min1+min2 <= money {
		return money - (min1 + min2)
	} else {
		return money
	}
}

// The solution that is better in memory.
/*func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	l, r := 0, len(prices) - 1
	res := -1
	for l < len(prices) && r >= 0 && l < r {
		total := money - prices[l] -prices[r]
		if total >= 0 {
			res = max(res,   total)

		}

		r--
	}
	if res < 0 {
		return money
	}
	return res
}

func max(a,b int) int {
	if a < b {
		return b
	}

	return a
}*/
