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

type Admin struct {
	Email    string
	Password string
	User
}

func (user *User) OutputUserDetails() {
	fmt.Println(user.FirstName, user.LastName, user.Birthdate)
}

func (user *User) ClearUserName() {
	(*user).FirstName = ""
	(*user).LastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
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

func NewAdmin(email, password string, appUser *User) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("missing email or password")
	}

	return &Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: (*appUser).FirstName,
			LastName:  (*appUser).LastName,
			Birthdate: (*appUser).Birthdate,
			CreatedAt: time.Now(),
		},
	}, nil
}
