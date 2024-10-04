package calculator

// CalculateProfit takes in the cost price, selling price, and discount percentage as input and returns the profit amount.
func CalculateProfit(
	revenue float64,
	expense float64,
	taxRate float64,
) (float64, float64, float64) {
	// Calculate Earning Before Tax (EBT)
	ebt := revenue - expense

	// Calculate Earning After Tax (profit)
	profit := ebt * (1 - taxRate/100)

	// Calculate Ratio (EBT/profit)
	var ratio float64
	if profit != 0 {
		ratio = ebt / profit
	} else {
		ratio = 0
	}

	return ebt, profit, ratio
}
