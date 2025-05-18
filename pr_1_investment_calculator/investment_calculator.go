package main

import (
	"fmt"
	"math"
)

func main() {
	// const inflationRate float64 = 6
	// var investmentAmount float64 = 1000
	// expectedReturnRate := 5.5
	// var years float64 = 10

	// var investmentAmount, years float64 = 1000, 10
	// investmentAmount, years, expectedReturnRate := 1000, 10, 5.5

	const inflationRate float64 = 6
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Println("Enter investment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter expected return rate:")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("Enter number of years:")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	var formattedFV string = fmt.Sprintf("Future value: %.2f\n", futureValue)
	var formattedFRV string = fmt.Sprintf("Future value adjusted for inflation: %.2f", futureRealValue)

	// fmt.Println("Future value: ", futureValue)
	// fmt.Println("Future value adjusted for inflation: ", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}
