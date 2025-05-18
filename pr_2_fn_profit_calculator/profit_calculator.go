package main

import "fmt"

func profitCalculator(purchasingPrice, sellingPrice float64) float64 {
	profit := sellingPrice - purchasingPrice
	return profit
}

func main() {
	var purchasingPrice, sellingPrice float64

	fmt.Print("Enter purchasing price: ")
	fmt.Scan(&purchasingPrice)

	fmt.Print("Enter selling price: ")
	fmt.Scan(&sellingPrice)

	profit := profitCalculator(purchasingPrice, sellingPrice)

	fmt.Printf("Profit: %.2f", profit)
}
