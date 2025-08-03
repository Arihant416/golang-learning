package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("one or more fields are empty")
	}
	return &User{
		firstName, lastName, birthDate, time.Now(),
	}, nil
}

func (person User) OutputUserDetails() {
	fmt.Println("FirstName", person.firstName)
	fmt.Println("LastName", person.lastName)
	fmt.Println("BirthDay", person.birthDate)
	fmt.Println("CreatedAt", person.createdAt)
}

func (person *User) ClearUserName() {
	person.firstName = "Curio"
	person.lastName = "Quest"
}
