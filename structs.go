package main

import (
	"fmt"

	"github.com/spwhitehead/structs-go/user"
)



func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")


	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OuputUserDetails()
	admin.ClearUserName()
	admin.OuputUserDetails()

	appUser.OuputUserDetails()
	appUser.ClearUserName()
	appUser.OuputUserDetails()

	// ... do something awesome with that gathered data!
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
