package main

import (
	"github.com/nearbyapaper/go_101/demo"
)

func main() {

	// var revenue float64
	// var expense float64
	// var taxRate float64

	// // Fetching User Data
	// fmt.Printf("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Printf("\nExpense: ")
	// fmt.Scan(&expense)

	// fmt.Printf("\nTax Rate: ")
	// fmt.Scan(&taxRate)

	// // Calculating Profit
	// ebt, profit, ratio := calculator.CalculateProfit(revenue, expense, taxRate)

	// // Use Backtrick for display
	// fmt.Printf(`EBT: %.2f
	// profit: %.2f
	// ratio: %.2f

	// ========================================
	// `, ebt, profit, ratio)

	// var investmentAmount float64
	// var timeInTears int
	// var expectedReturnRate float64
	// var inflationRate float64

	// fmt.Printf("InvestmentAmount :")
	// fmt.Scan(&investmentAmount)

	// fmt.Printf("\nTime in Years :")
	// fmt.Scan(&timeInTears)

	// fmt.Printf("\nExpected Return Rate :")
	// fmt.Scan(&expectedReturnRate)

	// fmt.Printf("\nInflation Rate :")
	// fmt.Scan(&inflationRate)

	// futureValue, realFutureValue := calculator.InvestmentCalculator(investmentAmount, timeInTears, expectedReturnRate, inflationRate)

	// sFutureValue := fmt.Sprintf("Future Value : %.2f", futureValue)
	// sRealFutureValue := fmt.Sprintf("\nReal Future Value : %.2f", realFutureValue)
	// fmt.Println(sFutureValue, sRealFutureValue)

	// fmt.Println("==========================================")
	demo.BankApp()
}
