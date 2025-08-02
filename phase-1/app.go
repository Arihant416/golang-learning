package main

import (
	"fmt"

	"example.com/phase1/basicmaths"
	"example.com/phase1/loops"
	maxofthree "example.com/phase1/maxOfThree"
	"example.com/phase1/numberchecker"
)

func main() {
	PrintWelcomeMessage()
	for {
		DisplayMenu()

		var choice int
		fmt.Print("Enter Choice: ")
		_, error := fmt.Scanln(&choice)
		if error != nil {
			fmt.Println("Invalid Choice Entered: Valid 1 to 6")
			continue
		}

		switch choice {
		case 1:
			basicmaths.PlayTheGameOfArithmetics()
		case 2:
			numberchecker.CheckNumberProperty()
		case 3:
			loops.PlayWithLoops()
		case 4:
			maxofthree.FindMaxOfThree()
		default:
			fmt.Println("Currently Not supported")
			fmt.Println("Exiting ...")
			return
		}
	}
}
