package model

type Aircraft struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Make  string `json:"make"`
	Model string `json:"model"`
}

func (a Aircraft) TableName() string {
	return "aircraft"
}

func ExampleAircraft() Aircraft {
	return Aircraft{
		ID:    1,
		Name:  "Snoopy",
		Year:  1967,
		Make:  "Cessna",
		Model: "172",
	}
}
