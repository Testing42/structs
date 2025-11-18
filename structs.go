package main

import (
	"fmt"

	"example.com/structs/user"
)

//defining your own custom types
//I f you put an uppercase on types it allows the function to be used in many fuctions and files

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//you can omit the firstname: same for other values but I would keep it as it could get confusing
	//you can omit a whole value
	var appUser *user.User

	appUser, err := user.NewUser(userfirstName, userlastName, userbirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
