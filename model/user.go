package model

type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Clubs     []*Club `gorm:"many2many:membership"`
}

func (u User) TableName() string {
	return "user"
}

func ExampleUser() User {
	exampleClub := ExampleClub()
	return User{
		ID:        1,
		FirstName: "Connor",
		LastName:  "Luckett",
		Clubs:     []*Club{&exampleClub},
	}
}
