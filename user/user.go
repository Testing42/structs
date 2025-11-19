package user

import (
	"errors"
	"fmt"
	"time"
)

// To expose variables in a struct the variable must be upper case
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

//You can make a new struct with all the information of another struct

type Admin struct {
	email    string
	password string
	User
}

/*
	in order to practice with pointers instead of copying data I will use a pointer so that

The same stored pointer data is used. Doing it this way saves data. This isn't needed
most of the time but helps with practice
*/

// this is how you add a stuct to a function to create methods
// you don't have to put a derefence when calling the struct but if you don't a copy of the data will be created.

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	//..

	/*fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)
	technically you need to dereference the pointer as yI have done above
	but golang offers the option to not do that with structs so the below works.
	*/
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// this shows how to edit struct values
// when editing a struct make sure to actually get the pointer by addin * infront of the pointer to dereference.
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// this is the constructor it makes the struct and can allow easy validation
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	//setting validation for struct
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
