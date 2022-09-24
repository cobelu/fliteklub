package model

type Ownership struct {
	ID         int      `json:"id"`
	ClubID     int      `gorm:"column:club_id"`
	Club       Club     `json:"club"`
	AircraftID int      `gorm:"column:aircraft_id"`
	Aircraft   Aircraft `json:"aircraft"`
}

func (o Ownership) TableName() string {
	return "ownership"
}

func ExampleOwnership() Ownership {
	return Ownership{
		ID:       1,
		Club:     ExampleClub(),
		Aircraft: ExampleAircraft(),
	}
}
