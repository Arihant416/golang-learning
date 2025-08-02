package main

import "fmt"

func FindFactorial() {
	fmt.Println("Hello, let's find factorial of a number")
	for {
		var choice int
		fmt.Println("Enter choice")
		fmt.Println("1. Find Factorial Iteratively")
		fmt.Println("2. Find Factorial Recursively")
		fmt.Println("Press any other key to exit")
		_, error := fmt.Scanln(&choice)
		if error != nil {
			fmt.Println("Invalid Choice")
			return
		}

		switch choice {
		case 1:
			findFactorialIteratively()
		case 2:
			findFactorialRecursively()
		default:
			fmt.Println("Other key pressed..Dropping and Exiting")
			return
		}
	}
}

func findFactorialIteratively() {
	var num int
	fmt.Print("Enter number: ")
	_, error := fmt.Scanln(&num)
	if error != nil {
		fmt.Println("Invalid number")
		return
	}
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	fmt.Println("Factorial of number ", num, " is ", result)
}

func findFactorialRecursively() {
	var num int
	fmt.Print("Enter number: ")
	_, error := fmt.Scanln(&num)
	if error != nil {
		fmt.Println("Invalid number")
		return
	}
	result := recursiveHelper(num)
	fmt.Println("Factorial of number ", num, " is ", result)
}

func recursiveHelper(num int) int {
	if num <= 1 {
		return 1
	}
	return num * recursiveHelper(num-1)
}
