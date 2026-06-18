package main

import "fmt"

type product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{1.9, 2.9}
	fmt.Println(prices[0:1])
	prices[1] = 3.9

	prices = append(prices, 4.9)
	fmt.Println(prices)
	prices = prices[1:]
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{1.9, 2.9, 3.9, 4.9}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
