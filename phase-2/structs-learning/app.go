package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var person *user.User
	person, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("Error while creating user -> ", err)
		return
	}

	person.OutputUserDetails()
	person.ClearUserName()
	person.OutputUserDetails()

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
	fmt.Scanln(&value)
	return value
}
