package calculator

// CalculateProfit takes in the cost price, selling price, and discount percentage as input and returns the profit amount.
func CalculateProfit(
	revenue float64,
	expense float64,
	taxRate float64,
) (float64, float64, float64) {
	// Calculate Earning Before Tax (EBT)
	ebt := calculateEBT(revenue, expense)

	// Calculate Earning After Tax (profit)
	profit := calculateProfit(ebt, taxRate)

	// Calculate Ratio (EBT/profit)
	ratio := calculateRatio(ebt, profit)

	return ebt, profit, ratio
}

func calculateEBT(revenue float64, expense float64) (ebt float64) {
	ebt = revenue - expense
	return
}

func calculateProfit(ebt float64, taxRate float64) (profit float64) {
	profit = ebt * (1 - taxRate/100)
	return profit
}

func calculateRatio(ebt float64, profit float64) (ratio float64) {
	if profit != 0 {
		ratio = ebt / profit
	} else {
		ratio = 0
	}
	return ratio
}
