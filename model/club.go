package model

type Club struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Location string      `json:"location"`
	Aircraft []*Aircraft `gorm:"many2many:ownership;"`
}

func (c Club) TableName() string {
	return "club"
}

func ExampleClub() Club {
	exampleAircraft := ExampleAircraft()
	return Club{
		ID:       1,
		Name:     "TAC",
		Location: "KGYI",
		Aircraft: []*Aircraft{&exampleAircraft},
	}
}
