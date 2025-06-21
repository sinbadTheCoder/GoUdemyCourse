package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years:")
	fmt.Scan(&years)

	futureValue := calcFutureValue(investmentAmount, expectedReturnRate, years)
	futureValueStr := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	futureValueAlt := calcFutureValueAlt(investmentAmount, expectedReturnRate, years)
	futureValueStr = fmt.Sprintf("Future Value: %.0f\n", futureValueAlt)

	futureRealVAlue := futureValue / math.Pow(1+inflationRate/100, years)
	adjFutureValueStr := fmt.Sprintf("Future Value (adj for inflation): %.0f\n", futureRealVAlue)

	fmt.Print(futureValueStr, adjFutureValueStr)
}

func calcFutureValue(investmentAmount, expectedReturnRate, years float64) float64 {
	return investmentAmount * math.Pow((1+expectedReturnRate/100), years)
}
func calcFutureValueAlt(investmentAmount, expectedReturnRate, years float64) (fv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	return
}
