package model

import (
	"fliteklub/misc"
	"time"
)

type Reservation struct {
	ID         int       `json:"id"`
	AircraftID int       `gorm:"column:aircraft_id"`
	Aircraft   Aircraft  `json:"aircraft"`
	UserID     int       `gorm:"column:user_id"`
	User       User      `json:"user"`
	Start      time.Time `gorm:"type:time" json:"start"`
	End        time.Time `gorm:"type:time" json:"end"`
}

func (r Reservation) TableName() string {
	return "reservations"
}

func ExampleReservation() Reservation {
	return Reservation{
		ID:       1,
		Aircraft: ExampleAircraft(),
		User:     ExampleUser(),
		Start:    misc.OneHourLater(),
		End:      misc.OneHourLater(),
	}
}
