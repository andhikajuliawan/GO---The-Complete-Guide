package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	transformed := transformNumber(&numbers, func(val int) int {
		return val * 2
	})

	fmt.Println(transformed)

	double := createTransformer(2)
	triple := createTransformer(3)

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

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
