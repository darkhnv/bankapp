package main

import (
	"fmt"
	"os"
	"strconv"

	"basics/bankapp/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			accountBalance = deposit(accountBalance)
		case 3:
			accountBalance = withdraw(accountBalance)
		case 4:
			exitBank(accountBalance)
			return
		default:
			fmt.Println("Invalid choice! Please select a valid option.")
		}
	}
}

func deposit(balance float64) float64 {
	fmt.Print("Enter the amount to deposit: ")
	amount := getAmount()

	if amount <= 0 {
		fmt.Println("Invalid amount. Deposit amount must be greater than 0.")
		return balance
	}

	balance += amount
	fmt.Println("Deposit successful! Your new balance is:", balance)
	fileops.WriteFloatToFile(accountBalanceFile, balance)
	return balance
}

func withdraw(balance float64) float64 {
	fmt.Print("Enter the amount to withdraw: ")
	amount := getAmount()

	if amount <= 0 {
		fmt.Println("Invalid amount. Withdrawal amount must be greater than 0.")
		return balance
	}

	if amount > balance {
		fmt.Println("Insufficient funds! You cannot withdraw more than your balance.")
		return balance
	}

	balance -= amount
	fmt.Println("Withdrawal successful! Your new balance is:", balance)
	fileops.WriteFloatToFile(accountBalanceFile, balance)
	return balance
}

func exitBank(balance float64) {
	fmt.Println("Thank you for banking with us!")
	fmt.Println("Your final balance is:", balance)
	fmt.Println("Goodbye!")
}

func getAmount() float64 {
	var input string
	fmt.Scanln(&input)
	amount, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return 0
	}
	return amount
}

func presentOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
