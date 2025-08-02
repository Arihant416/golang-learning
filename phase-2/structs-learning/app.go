package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (person user) outputUserDetails() {
	fmt.Println("FirstName", person.firstName)
	fmt.Println("LastName", person.lastName)
	fmt.Println("BirthDay", person.birthDate)
	fmt.Println("CreatedAt", person.createdAt)
}

func (person *user) clearUserName() {
	person.firstName = "Curio"
	person.lastName = "Quest"
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var person user = user{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}

	person.outputUserDetails()
	person.clearUserName()
	person.outputUserDetails()

	// fmt.Println(firstName, lastName, birthDate)
}

// func outputUserDetails(person *user) {
// 	fmt.Println("FirstName", person.firstName)
// 	fmt.Println("LastName", person.lastName)
// 	fmt.Println("BirthDay", person.birthDate)
// 	fmt.Println("CreatedAt", person.createdAt)

// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
