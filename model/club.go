package model

type Club struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (c Club) TableName() string {
	return "club"
}

func ExampleClub() Club {
	return Club{
		ID:       1,
		Name:     "TAC",
		Location: "KGYI",
	}
}
