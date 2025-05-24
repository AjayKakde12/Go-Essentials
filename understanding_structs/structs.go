package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (user User) outputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

func (user *User) clearUserName() {
	(*user).firstName = ""
	(*user).lastName = ""
}

func outputUserDetailsUsingPointer(user *User) {
	fmt.Println((*user).firstName, (*user).lastName, (*user).birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func newUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("you should have entered first name, last name & birth date")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(firstName, lastName, birthdate)

	if err != nil {
		panic(err)
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	outputUserDetailsUsingPointer(appUser)

	appUser.clearUserName()
	appUser.outputUserDetails()
}
