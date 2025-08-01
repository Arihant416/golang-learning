package main

import (
	"fmt"

	"example.com/phase1/basicmaths"
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
		default:
			fmt.Println("Currently Not supported")
			fmt.Println("Exiting ...")
			return
		}
	}
}
