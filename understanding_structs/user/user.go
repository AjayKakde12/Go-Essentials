package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

func (user *User) OutputUserDetails() {
	fmt.Println(user.FirstName, user.LastName, user.Birthdate)
}

func (user *User) ClearUserName() {
	(*user).FirstName = ""
	(*user).LastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("you should have entered first name, last name & birth date")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}
