package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(first, last, birthdate string) (*User, error) {
	if first == "" || last == "" || birthdate == "" {
		return nil, errors.New("invalid data entered")
	}

	var u = &User{
		firstName: first,
		lastName:  last,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
	return u, nil
}

func (u User) OutputUserDetails() {
	//shows shorthand to dereference pointers to structures
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

// modifying the value of a struct requires pointer here
func (u *User) ClearUserName() {
	(*u).firstName = ""
	u.lastName = ""
}
