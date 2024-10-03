package main // The Importanf of The name main <-- "main" is tell the go project this is starter of application

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount = 30000
	// var expectedReturnenRate = 5.5
	// var investmentYears = 5

	// short way to declare variable use ':=' operator GO will assign type that should be to the variable
	investmentAmount := 30000.0
	expectedReturnenRate := 5.5
	investmentYears := 5.0

	// calculate future value
	// math.Pow(base, exponent) returns base raised to exponent
	// math.Pow(1+expectedReturnenRate/100, float64(investmentYears)) calculates the future value for each year
	// float64(investmentAmount) converts the investment amount to float64 for the calculation
	// * is used to perform multiplication in Go, not ^ for exponentiation
	// float64 is necessary to perform mathematical operations on numbers of different types (int and float64) in Go
	// float64(investmentAmount) converts investmentAmount to float64 for the calculation
	// * is used to perform multiplication in Go, not ^ for exponentiation
	// float64 is necessary to perform mathematical operations

	var futureValue = investmentAmount * math.Pow((1+expectedReturnenRate/100), investmentYears)
	// convert investmentAmount to float64 by float64(number: int)
	// can't mix between different type for operation

	fmt.Printf("The future value of investment after %.2f years is: $%.2f\n", investmentYears, futureValue)
}
