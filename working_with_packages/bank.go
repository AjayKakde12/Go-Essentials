package main

import (
	"fmt"

	"github.com/AjayKakde12/Go-Essentials/working-with-packages/fileops"
)

const BALANCE_FILE string = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(BALANCE_FILE)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------------")
		panic("Time to panic.............")
	}
	var choice int

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, BALANCE_FILE)
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
			fileops.WriteFloatToFile(accountBalance, BALANCE_FILE)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for chosing our bank!")
}
