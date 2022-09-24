package model

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u User) TableName() string {
	return "user"
}

func ExampleUser() User {
	return User{
		ID:        1,
		FirstName: "Connor",
		LastName:  "Luckett",
	}
}
