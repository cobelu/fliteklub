package model

type Club struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Location string      `json:"location"`
	Aircraft []*Aircraft `gorm:"many2many:ownerships;"`
	Members  []*User     `gorm:"many2many:memberships;"`
}

func (c Club) TableName() string {
	return "clubs"
}

func (c Club) ModelName() string {
	return "Club"
}

func ExampleClub() Club {
	exampleAircraft := ExampleAircraft()
	exampleMember := ExampleUser()
	return Club{
		ID:       1,
		Name:     "TAC",
		Location: "KGYI",
		Aircraft: []*Aircraft{&exampleAircraft},
		Members:  []*User{&exampleMember},
	}
}
