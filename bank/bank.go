package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	return fileops.GetFloatFromFile(accountBalanceFile)
}

func writeBalanceToFile(balance float64) {
	fileops.WriteFloatToFile(accountBalanceFile, balance)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		panic("Unable to retrieve balance. Exiting.")
	}
	fmt.Println("Welcome go Go Bank!")
	fmt.Println("Reach us at ", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("How much would you like to deposit? ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New Amount: ", accountBalance)
		case 3:
			var amount float64
			fmt.Print("How much would you like to withdraw? ")
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			if amount > accountBalance {
				fmt.Printf("You cannot withraw greater than your balance of %.2f\n", accountBalance)
				continue
			}

			accountBalance -= amount
			writeBalanceToFile((accountBalance))
			fmt.Println("Balance updated! New Amount: ", accountBalance)
		default:
			fmt.Println("Goodbye.")
			fmt.Println("Thanks for choosing our bank.")
			return
		}
	}
}
