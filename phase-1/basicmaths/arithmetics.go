package basicmaths

import (
	"fmt"
)

func displayMenu() {
	fmt.Println("\nSelect an Operation from the Menu")
	fmt.Println("1. Addition of two numbers")
	fmt.Println("2. Subtraction of two numbers")
	fmt.Println("3. Multiplication of two numbers")
	fmt.Println("4. Division")
	fmt.Println("5. Modulo")
	fmt.Println("6. Exit the Game!!!")
}

func PlayTheGameOfArithmetics() {
	for {
		displayMenu()
		var choice int
		fmt.Print("Enter choice: ")
		_, error := fmt.Scanln(&choice)
		if error != nil {
			fmt.Println()
			fmt.Println("Invalid operation entered...")
			continue
		}

		if choice == 6 {
			fmt.Println("You've opted to exit the Game!! Sayonara....")
			return
		}

		var num1, num2 int
		fmt.Print("Enter N1: ")
		_, err := fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("Invalid character entered..Expected a number")
			return
		}
		fmt.Print("Enter N2: ")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("Invalid character entered. Expected a number")
			return
		}

		switch choice {
		case 1:
			addTwoNumbers(num1, num2)
		case 2:
			subtractTwoNumbers(num1, num2)
		case 3:
			multiplyTwoNumbers(num1, num2)
		case 4:
			divideANumberByAnotherNumber(num1, num2)
		case 5:
			showModulusOperation(num1, num2)
		default:
			fmt.Println("Press 6 to exit..else please enter valid key")
			continue
		}
	}
}

func addTwoNumbers(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println("Result of Addition:", result)
}

func subtractTwoNumbers(num1 int, num2 int) {
	result := num1 - num2
	fmt.Println("Result of Subtraction:", result)
}

func multiplyTwoNumbers(num1 int, num2 int) {
	result := num1 * num2
	fmt.Println("Result of Multiplication:", result)
}

func divideANumberByAnotherNumber(num1 int, num2 int) {
	result := num1 / num2
	fmt.Println("Result of Division", result)
}

func showModulusOperation(num1 int, num2 int) {
	result := num1 % num2
	fmt.Println("Result of Modulo", result)
}
