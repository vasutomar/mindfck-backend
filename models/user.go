package models

import "github.com/google/uuid"

type User struct {
	FirstName string
	LastName  string
	UserId    int
}

func (u User) createNewUser(f, l string) User {
	user := User{
		FirstName: f,
		LastName:  l,
		UserId:    uuid.ClockSequence(),
	}
	return user
}
