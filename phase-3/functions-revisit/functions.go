package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubledManual := doubleNumbers(&numbers)
	doubled := transform(&numbers, double)
	tripled := transform(&numbers, triple)

	fmt.Println(doubledManual)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func doubleNumbers(numbers *[]int) []int {
	doubled := []int{}
	for _, num := range *numbers {
		doubled = append(doubled, double(num))
	}
	return doubled
}

func transform(numbers *[]int, transformer transformFn) []int {
	transformed := []int{}
	for _, num := range *numbers {
		transformed = append(transformed, transformer(num))
	}
	return transformed
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFn() transformFn {
	return double
}
