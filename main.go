package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumber(numbers *[]int, transform transformFn) []int {
	result := []int{}

	for _, val := range *numbers {
		result = append(result, transform(val))
	}

	return result
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
