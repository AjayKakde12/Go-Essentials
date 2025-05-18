package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 6

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Println("Enter investment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter expected return rate:")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("Enter number of years:")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future value adjusted for inflation: %.2f", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
