package model

import (
	"fliteklub/misc"
)

type Club struct {
	misc.GormModel
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Created  misc.Date `json:"created"`
}

func ExampleClub() Club {
	return Club{
		GormModel: misc.GormModel{},
		Name:      "TAC",
		Location:  "KGYI",
		Created:   misc.Today(),
	}
}
