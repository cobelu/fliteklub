package model

import (
	"fliteklub/misc"
)

type User struct {
	misc.GormModel
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Created   misc.Date `json:"created"`
}

func ExampleUser() User {
	return User{
		GormModel: misc.GormModel{},
		FirstName: "Connor",
		LastName:  "Luckett",
		Created:   misc.Today(),
	}
}
