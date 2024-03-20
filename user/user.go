package user

import (
	"errors"
	"fmt"
	"time"
)


type User struct {		// The same as a class in python
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}		

func (u *User) OuputUserDetails() {
	// ...
	
	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	// validation checks
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First and last names and birthdate are required.")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}