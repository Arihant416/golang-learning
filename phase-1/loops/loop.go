package loops

import "fmt"

func displayMenu() {
	fmt.Println("You can do the following")
	fmt.Println("1. Print 1 to N")
	fmt.Println("2. Print a Countdown")
	fmt.Println("3. Print a Pattern")
	fmt.Println("Any other key to Exit loop")
}

func PlayWithLoops() {
	fmt.Println("\nHello there.. Let's play with loops")
	for {
		displayMenu()
		var choice int
		fmt.Print("Enter choice: ")
		_, error := fmt.Scanln(&choice)

		if error != nil {
			fmt.Println("Entered an exit choice")
			return
		}

		switch choice {
		case 1:
			print1ToN()
		case 2:
			printCountDown()
		case 3:
			printAPattern()
		default:
			fmt.Println("Exiting loop.....")
			return
		}
	}
}

func print1ToN() {
	var N int
	fmt.Print("\nEnter N: ")
	_, error := fmt.Scanln(&N)
	if error != nil {
		fmt.Println("Invalid N input")
		return
	}
	for i := 1; i <= N; i++ {
		fmt.Println(i)
	}
}

func printCountDown() {
	var N int
	fmt.Print("\nEnter N: ")
	_, error := fmt.Scanln(&N)
	if error != nil {
		fmt.Println("Invalid N Input")
		return
	}
	for i := N; i >= 0; i-- {
		fmt.Println(i)
	}
}

func printAPattern() {
	var totalLines int
	fmt.Print("Enter Depth: ")
	_, err := fmt.Scanln(&totalLines)
	if err != nil {
		fmt.Println("Invalid Depth Given")
		return
	}
	spaceCount := totalLines - 1
	for i := 1; i <= totalLines; i++ {
		for j := 1; j <= spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := totalLines - spaceCount; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
		spaceCount--
	}
}
