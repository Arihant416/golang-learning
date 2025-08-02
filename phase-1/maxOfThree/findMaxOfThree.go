package maxofthree

import "fmt"

func FindMaxOfThree() {
	fmt.Println("Hello, let's find maximum of three numbers")
	for {
		var a, b, c int
		fmt.Print("Enter a: ")
		_, error := fmt.Scanln(&a)
		if error != nil {
			fmt.Println("Input Invalid")
			continue
		}
		fmt.Print("Enter b: ")
		_, error = fmt.Scanln(&b)
		if error != nil {
			fmt.Println("Input Invalid")
			continue
		}
		fmt.Print("Enter c: ")
		_, error = fmt.Scanln(&c)
		if error != nil {
			fmt.Println("Input Invalid")
			continue
		}

		result := helperFindMaxOfThree(a, b, c)
		fmt.Println("Max is: ", result)
		return
	}
}

func helperFindMaxOfThree(a int, b int, c int) int {
	if a > b {
		if b > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
