package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your (MM/DD/YYYY): ")

	var appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)
}

func outputUserDetails(userData *user) {
	//shows shorthand to dereference pointers to structures
	fmt.Println((*userData).firstName, userData.lastName, userData.birthdate, userData.createdAt)
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
