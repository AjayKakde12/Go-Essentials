package main

import (
	"fmt"

	"github.com/AjayKakde12/Go-Essentials/structs/user"
)

func outputUserDetailsUsingPointer(u *user.User, admin *user.Admin) {
	fmt.Println((*u).FirstName, (*u).LastName, (*u).Birthdate, (*admin).Email, (*admin).Password)
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
	email := getUserData("Please enter your email: ")
	password := getUserData("Please enter password")

	appUser, errUser := user.New(firstName, lastName, birthdate)
	if errUser != nil {
		panic(errUser)
	}

	admin, errAdmin := user.NewAdmin(email, password, appUser)
	if errAdmin != nil {
		panic(errAdmin)
	}

	// ... do something awesome with that gathered data!

	admin.OutputUserDetails()
	outputUserDetailsUsingPointer(appUser, admin)

	admin.ClearUserName()
	admin.OutputUserDetails()
}
