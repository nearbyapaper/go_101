package calculator

import "math"

func InvestmentCalculator(
	investmentAmount float64,
	timeInYears int,
	expectedReturnRate float64,
	inflationRate float64,
) (float64, float64) {
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(timeInYears))
	// realFutureValue := futureValue / math.Pow((1+inflationRate/100), float64(timeInYears))

	futureValue, realFutureValue := calculatorFutureValues(investmentAmount, expectedReturnRate, timeInYears, inflationRate)

	return futureValue, realFutureValue
}

func calculatorFutureValues(investmentAmount float64, expectedReturnRate float64, timeInYears int, inflationRate float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(timeInYears))
	rfv := fv / math.Pow((1+inflationRate/100), float64(timeInYears))

	return fv, rfv
}
