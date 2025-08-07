package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.9, 11.9}
	fmt.Println(prices[1])
	prices[1] = 9.99
	fmt.Println(prices)
	// prices[2] = 10 This will  cause an Index Out of range error
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 15.99, 21.99, 22.30}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)
// }

// Array are data structures holding different values of the same type
