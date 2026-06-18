package main

import "fmt"

type product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{1.9, 2.9, 3.9, 4.9}
	fmt.Println(prices)

	productNames[2] = "A Carpet"
	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
}
