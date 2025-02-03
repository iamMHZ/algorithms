package richest_customer_wealth

func MaximumWealth(accounts [][]int) int {
	var max int = 0
	for i := 0; i < len(accounts); i++ {
		var customer_wealth int = 0
		for j := 0; j < len(accounts[i]); j++ {
			customer_wealth += accounts[i][j]
		}

		if customer_wealth > max {
			max = customer_wealth
		}

	}

	return max

}
