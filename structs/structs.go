package main

import (
	"fmt"
	"reflect"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func newUser(first, last, birthdate string) *user {
	return &user{
		firstName: first,
		lastName:  last,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func (u user) outputUserDetails() {
	//shows shorthand to dereference pointers to structures
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

// modifying the value of a struct requires pointer here
func (u *user) clearUserName() {
	(*u).firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your (MM/DD/YYYY): ")

	var appUser = *newUser(userFirstName, userLastName, userBirthdate)
	fmt.Printf("appUser is of type [%s]\n", reflect.ValueOf(appUser).Kind())

	var appUser2 = newUser(userFirstName, userLastName, userBirthdate)
	fmt.Printf("appUser2 is of type [%s]\n", reflect.ValueOf(appUser2).Kind())

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
