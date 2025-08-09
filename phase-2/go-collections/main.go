package main

import "fmt"

func main() {
	userNames := make([]string, 2) // Make creates an element of size 2
	userNames[0] = "Arihant"
	userNames = append(userNames, "Max")    // Now append creates a new array with max on the end after 2 empty spaces
	userNames = append(userNames, "Manuel") // Same , which makes array of length 4

	fmt.Println(userNames)

	courseRatings := map[string]float64{}
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	fmt.Println(courseRatings)
}
