package model

import (
	"fliteklub/misc"
)

type User struct {
	misc.GormModel
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Created   string `json:"created"`
}

type Club struct {
	misc.GormModel
	Name     string `json:"name"`
	Location string `json:"location"`
	Created  string `json:"created"`
}

type Aircraft struct {
	misc.GormModel
	Name     string `json:"name"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Location string `json:"location"`
}
