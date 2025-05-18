package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const BALANCE_FILE string = "balance.txt"

func writeBalanceToFile(balance float64) {
	balancetext := fmt.Sprint(balance)
	os.WriteFile(BALANCE_FILE, []byte(balancetext), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(BALANCE_FILE)
	if err != nil {
		return 0, errors.New("storage not found")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("storage contains invalid data")
	}
	return balance, nil
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------------")
		panic("Time to panic.............")
	}
	var choice int

	for {
		fmt.Println("Welcome to Go Bank")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		fmt.Printf("Selected: %v\n", choice)

		if choice == 1 {
			fmt.Printf("Your account balance is %.2f\n", accountBalance)
		} else if choice == 2 {
			var depositeAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositeAmount)
			if depositeAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += depositeAmount
			fmt.Println("Balance updated: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			var withdrawMoney float64
			fmt.Print("Your withdraw amount: ")
			fmt.Scan(&withdrawMoney)
			if withdrawMoney > accountBalance || withdrawMoney <= 0 {
				fmt.Println("Invalid amount. Must be less than account balancd and greater than 0")
				continue
			}
			accountBalance -= withdrawMoney
			fmt.Println("Balance updated: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for chosing our bank!")
}
