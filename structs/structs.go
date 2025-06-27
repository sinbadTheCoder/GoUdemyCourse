package main

import (
	"fmt"
	"reflect"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your (MM/DD/YYYY): ")

	var appUser *user.User
	var err error
	appUser, err = user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("appUser is of type [%s]\n", reflect.ValueOf(appUser).Kind())

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
