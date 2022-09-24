package model

type Club struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	LocationID int         `gorm:"column:user_id"`
	Location   Airport     `json:"location"`
	Aircraft   []*Aircraft `gorm:"many2many:ownerships;"`
	Members    []*User     `gorm:"many2many:memberships;"`
}

func (c Club) TableName() string {
	return "clubs"
}

func ExampleClub() Club {
	exampleAircraft := ExampleAircraft()
	exampleMember := ExampleUser()
	exampleAirport := ExampleAirport()
	return Club{
		ID:       1,
		Name:     "TAC",
		Location: exampleAirport,
		Aircraft: []*Aircraft{&exampleAircraft},
		Members:  []*User{&exampleMember},
	}
}
