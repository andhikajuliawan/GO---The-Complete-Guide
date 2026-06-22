package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	// userName := []string{}
	// userName := make([]string, 2)
	userName := make([]string, 2, 5)
	userName[0] = "andhika"

	userName = append(userName, "max")
	userName = append(userName, "william")

	fmt.Println(userName)

	// courseRatings := map[string]float64{}
	courseRatings := make(floatMap, 3)

	courseRatings["Go"] = 7.8
	courseRatings["JS"] = 4.6
	courseRatings["Angular"] = 4.7

	courseRatings.output()
	// fmt.Println(courseRatings)
}
