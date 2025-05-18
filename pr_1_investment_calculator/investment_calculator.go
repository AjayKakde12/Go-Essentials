package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	// var investmentAmount, years float64 = 1000, 10
	// investmentAmount, years, expectedReturnRate := 1000, 10, 5.5

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
