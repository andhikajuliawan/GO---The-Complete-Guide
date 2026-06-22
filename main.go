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
}
