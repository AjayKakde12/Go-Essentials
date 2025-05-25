package main

import "fmt"

func main() {
	productNames := [4]string{"A Book"}

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "Apartments"

	fmt.Println(productNames)

	fmt.Println(prices[1])

	featuredPrices := prices[1:3] //Slices in Go startIndex: endIndex + 1
	fmt.Println(featuredPrices)
}
