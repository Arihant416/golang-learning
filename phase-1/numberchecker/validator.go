package numberchecker

import "fmt"

func CheckNumberProperty() {
	fmt.Println("\nWelcome to Number Checker")
	for {
		var number int
		fmt.Print("\nEnter number: ")
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Character not a number\nExiting Checker Function")
			return
		}

		analyseProperty(&number)
	}

}

func analyseProperty(n *int) {
	if *n == 0 {
		fmt.Println("Number is 0")
	} else if *n%2 == 0 {
		if *n < 0 {
			fmt.Println("Number is even and Negative")
		} else {
			fmt.Println("Number is even and Positive")
		}
	} else {
		if *n > 0 {
			fmt.Println("Number is odd and Positive")
		} else {
			fmt.Println("Number is odd and Negative")
		}
	}
}
