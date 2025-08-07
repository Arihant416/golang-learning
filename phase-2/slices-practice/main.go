package main

import "fmt"

type Product struct {
	title  string
	id     string
	prices float64
}

func main() {
	hobbies := []string{"Coding", "Singing", "Walking"}
	fmt.Println(hobbies)
	fmt.Println("First Hobby: ", hobbies[0])
	fmt.Println(hobbies[1:])

	slice1 := hobbies[:2]
	slice2 := hobbies[0:2]
	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice2: ", slice2)

	slice3 := slice1[1:3]
	fmt.Println(slice3)

	goals := []string{"Tech lead", "Parent"}
	goals[1] = "Learn Go"
	goals = append(goals, "Be a good dad")
	fmt.Println(goals)

	products := []Product{
		{
			"first-product",
			"A First Product",
			12.99,
		},
		{
			"second-product",
			"A Second Product",
			129.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A Third Product",
		15.99,
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
