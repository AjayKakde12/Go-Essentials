package main

import (
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
	fmt.Scan(&value)
	return value
}

func main() {
	appUser := User{
		firstName: getUserData("Please enter your first name: "),
		lastName:  getUserData("Please enter your last name: "),
		birthdate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	outputUserDetailsUsingPointer(&appUser)

	appUser.clearUserName()
	appUser.outputUserDetails()
}
