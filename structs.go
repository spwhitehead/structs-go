package main

import (
	"fmt"
	"time"
)

type user struct {		// The same as a class in python
	firstName string
	lastName string
	birthDate string
	age int
	createdAt time.Time
}		

func (u *user) ouputUserDetails() {
	// ...
	
	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")


	var appUser *user

	appUser = newUser(userFirstName, userLastName, userBirthDate)
	appUser.ouputUserDetails()
	appUser.clearUserName()
	appUser.ouputUserDetails()

	// ... do something awesome with that gathered data!
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
