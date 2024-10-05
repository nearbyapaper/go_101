package demo

import (
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func BankApp() {
	fmt.Println("Welcome to the Banking Application!")
	balance := getBalanceFromFile()
	writeBalanceToFile(balance)

	for {
		fmt.Println("\nPlease choose an option:")
		fmt.Println("1. Deposit Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			balance = depositMoney(balance)
			continue

		case 2:
			balance = withdrawMoney(balance)
			continue

		case 3:
			checkBalance(balance)
			continue

		case 4:
			fmt.Println("Thank you for using the Banking Application!")
			fmt.Println("================================================")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
			continue
		}

	}

}

func writeBalanceToFile(balance float64) {
	data := []byte(fmt.Sprintf("%.2f", balance)) // change to bytes format for write
	os.WriteFile(fileName, data, 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(fileName)
	// _ to tell the program that I know will have data return but I'm not interesting
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func checkBalance(balance float64) {
	fmt.Printf("\nYour balance : %.2f", balance)
}

func withdrawMoney(balance float64) float64 {
	fmt.Println("Please enter your withdraw amount :")
	var amount float64
	fmt.Scan(&amount)

	if amount > balance {
		fmt.Println("Insufficient balance.")
		return balance
	} else {
		balance = balance - amount
		fmt.Printf("\nWithdrawal successful. New balance is : %.2f", balance)
		writeBalanceToFile(balance)
		return balance
	}
}

func depositMoney(balance float64) float64 {
	fmt.Println("Please enter your deposit amount :")
	var amount float64
	fmt.Scan(&amount)

	if amount > 0 {
		balance = balance + amount
		fmt.Printf("\nDeposit successful. New balance is : %.2f", balance)
		writeBalanceToFile(balance)
		return balance
	} else {
		fmt.Println("\nInvalid amount. Please enter a positive number.")
		return balance
	}
}
