package calculator

import "math"

func InvestmentCalculator(
	investmentAmount float64,
	timeInYears int64,
	expectedReturnRate float64,
	inflationRate float64,
) (float64, float64) {
	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(timeInYears))
	realFutureValue := futureValue / math.Pow((1+inflationRate/100), float64(timeInYears))

	return futureValue, realFutureValue
}
