package main

import "fmt"

func main() {

	// userName := []string{}
	// userName := make([]string, 2)
	userName := make([]string, 2, 5)
	userName[0] = "andhika"

	userName = append(userName, "max")
	userName = append(userName, "william")

	fmt.Println(userName)

	// courseRatings := map[string]float64{}
	courseRatings := make(map[string]float64, 3)

	courseRatings["Go"] = 7.8
	courseRatings["JS"] = 4.6
	courseRatings["Angular"] = 4.7

	fmt.Println(courseRatings)

}
