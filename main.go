package main

import (
	"fmt"
	"math/rand"
)

type product struct {
	id    int
	title string
	price int
}

func main() {
	// 1
	hobbies := [3]string{"Badminton", "Swimming", "Tenis"}
	fmt.Println(hobbies)
	fmt.Println("-------------")

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	fmt.Println("-------------")

	// 3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)
	fmt.Println("-------------")

	// 4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)
	fmt.Println("-------------")

	// 5
	courseGoals := []string{"learn all the basic", "learn Go"}
	fmt.Println(courseGoals)
	fmt.Println("-------------")

	// 6
	courseGoals[1] = "Master in Go"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "learn microservices")
	fmt.Println(courseGoals)
	fmt.Println("-------------")

	// 7
	products := []product{
		{
			id:    rand.Intn(9999),
			title: "Shampo",
			price: 20000,
		}, {
			id:    rand.Intn(9999),
			title: "Soap",
			price: 5000,
		},
	}
	fmt.Println(products)

	newProduct := product{
			id:    rand.Intn(9999),
			title: "tooth paste",
			price: 10000,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
