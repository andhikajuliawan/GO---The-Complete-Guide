package main

import "fmt"

type product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := [4]float64{1.9, 2.9, 3.9, 4.9}
	fmt.Println(prices)
}
