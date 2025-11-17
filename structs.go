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

	//you can omit the firstname: same for other values but I would keep it as it could get confusing
	//you can omit a whole value
	var appUser = user{
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userbirthDate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

/*
	in order to practice with pointers instead of copying data I will use a pointer so that

The same stored pointer data is used. Doing it this way saves data. This isn't needed
most of the time but helps with practice
*/

// this is how you add a stuct to a function to create methods
// you don't have to put a derefence when calling the struct but if you don't a copy of the data will be created.
func (u *user) outputUserDetails() {
	//..

	/*fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)
	technically you need to dereference the pointer as yI have done above
	but golang offers the option to not do that with structs so the below works.
	*/
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// this shows how to edit struct values
// when editing a struct make sure to actually get the pointer by addin * infront of the pointer to dereference.
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
