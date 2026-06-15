package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func NewUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate are required.")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u *user) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func (u *user) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}
