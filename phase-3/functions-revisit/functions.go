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

	fmt.Println("Iterative Factorial of 5", findFactorialIteratively(5))
	fmt.Println("Recursive Factorial of 6", findFactorialRecursively(6))
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

func findFactorialIteratively(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func findFactorialRecursively(number int) int {
	if number <= 1 {
		return 1
	}
	return number * findFactorialRecursively(number-1)
}
