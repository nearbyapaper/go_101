package calculator

import (
	"errors"
	"fmt"
	"os"
)

const fileName = "profit.txt"

// CalculateProfit takes in the cost price, selling price, and discount percentage as input and returns the profit amount.
func CalculateProfit() {
	data, err := readFromFile()
	if err != nil {
		fmt.Println("Error reading profit from file:", err)
	} else {
		fmt.Println("Previous data from file:\n", data)
	}

	revenue := getUserInput("Input Your revenue: ")
	expense := getUserInput("Input Your expense: ")
	taxRate := getUserInput("Input Your tax rate: ")

	// Calculate Earning Before Tax (EBT)
	ebt := calculateEBT(revenue, expense)

	// Calculate Earning After Tax (profit)
	profit := calculateProfit(ebt, taxRate)

	// Calculate Ratio (EBT/profit)
	ratio := calculateRatio(ebt, profit)

	writeProfitToFile(ebt, profit, ratio)
}

func writeProfitToFile(ebt, profit, ratio float64) {
	data := []byte(fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio))
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func readFromFile() (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.New("can't read data from file")
	}

	return string(data), nil
}

func getUserInput(prompt string) float64 {
	var input float64
	for {
		fmt.Printf("%s", prompt)
		fmt.Scan(&input)
		if input > 0 {
			break
		}
		fmt.Println("Invalid input. Please enter a positive number.")
	}
	return input
}

func calculateEBT(revenue, expense float64) float64 {
	return revenue - expense
}

func calculateProfit(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calculateRatio(ebt, profit float64) float64 {
	if profit != 0 {
		return ebt / profit
	}
	return 0
}
