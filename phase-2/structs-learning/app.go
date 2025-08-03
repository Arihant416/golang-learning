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

	admin := user.NewAdmin("test@gmail.com", "knsdfoeinq")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	person.OutputUserDetails()
	person.ClearUserName()
	person.OutputUserDetails()

	// fmt.Println(firstName, lastName, birthDate)
}

// func outputUserDetails(person *user) {
// 	fmt.Println("FirstName", person.firstName)

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
