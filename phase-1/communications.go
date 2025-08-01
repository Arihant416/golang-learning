// Same Package different script!! To understand exported functions and unexported ones
package main

import "fmt"

func PrintWelcomeMessage(){
	fmt.Println("Welcome to GoPlayGround!!!")
}

func DisplayMenu(){
	fmt.Println("What'd you like to do?")
	fmt.Println("1. The Game of Arithmetics")
	fmt.Println("2. Check the properties of Number")
	fmt.Println("3. Play with loops")
	fmt.Println("4. Play Max of Three")
	fmt.Println("5. Factorial Game")
	fmt.Println("6. Understand Working of Pointers in Go!!!")
}