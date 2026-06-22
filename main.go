package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{5, 1, 3, 4}
	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&moreNumbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumber(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumber(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func transformNumber(numbers *[]int, transform transformFn) []int {
	result := []int{}

	for _, val := range *numbers {
		result = append(result, transform(val))
	}

	return result
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
