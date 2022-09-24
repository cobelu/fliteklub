package model

import (
	"fliteklub/misc"
)

type Aircraft struct {
	misc.GormModel
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Make  string `json:"make"`
	Model string `json:"model"`
}

func ExampleAircraft() Aircraft {
	return Aircraft{
		GormModel: misc.GormModel{},
		Name:      "Snoopy",
		Year:      1967,
		Make:      "Cessna",
		Model:     "172",
	}
}
