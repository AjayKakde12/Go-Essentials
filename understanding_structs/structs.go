package main

import (
	"fmt"

	"github.com/AjayKakde12/Go-Essentials/structs/user"
)

func outputUserDetailsUsingPointer(u *user.User) {
	fmt.Println((*u).FirstName, (*u).LastName, (*u).Birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		panic(err)
	}

	// ... do something awesome with that gathered data!

	(*appUser).OutputUserDetails()
	outputUserDetailsUsingPointer(appUser)

	(*appUser).ClearUserName()
	(*appUser).OutputUserDetails()
}
