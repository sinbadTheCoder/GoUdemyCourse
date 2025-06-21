package main

import (
	"errors"
	"fmt"
	"os"
)

var revenue, expenses, taxRate float64

func acceptFloatInput(text string) (value float64) {
	fmt.Println("Enter", text)
	numArgs, err := fmt.Scan(&value)
	if numArgs != 1 || err != nil {
		fmt.Println(err)
		panic("Illegal input")
	}
	err = validateInput(value)
	if err != nil {
		fmt.Println(err)
		panic("Invalid input")
	}
	return value
}

func validateInput(val float64) error {
	if val <= 0 {
		return errors.New("value is not greater than zero")
	}
	return nil
}

func main() {
	revenue = acceptFloatInput("revenue: ")
	expenses = acceptFloatInput("expenses: ")
	taxRate = acceptFloatInput("tax rate: ")

	ebt, profit, ratio := runCalculations(revenue, expenses, taxRate)

	outStr := fmt.Sprintf("EBT: %.2f, PROFIT: %.2f, RATIO: %.1f", ebt, profit, ratio)
	os.WriteFile("out.txt", []byte(outStr), 0644)

	fmt.Println("EBT: ", ebt, "PROFIT: ", profit, "RATIO:", ratio)
}

func runCalculations(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	//EBT=Revenue−Operating Expenses−Interest Expense
	netIncome := revenue - expenses
	taxAmt := taxRate / 100 * revenue
	ebt = netIncome
	profit = netIncome - taxAmt
	ratio = ebt / profit
	return ebt, profit, ratio
}
