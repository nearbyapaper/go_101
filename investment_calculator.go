package main // The Importanf of The name main <-- "main" is tell the go project this is starter of application

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello world!")
	var investmentAmount = 30000
	var expectedReturnenRate = 5.5
	var investmentYears = 5

	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnenRate/100), float64(investmentYears))
	// convert investmentAmount to float64 by float64(number: int)
	// can't mix between different type for operation

	fmt.Printf("The future value of investment after %d years is: $%.2f\n", investmentYears, futureValue)
}
