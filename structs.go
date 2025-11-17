package main

import (
	"fmt"
	"time"
)

//defining your own custom types
//I f you put an uppercase on types it allows the function to be used in many fuctions and files

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	//you can omit the firstname: same for other values but I would keep it as it could get confusing
	//you can omit a whole value
	appUser = user{
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userbirthDate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	outputUserDetails(appUser)
}

func outputUserDetails(u user) {
	//..
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
