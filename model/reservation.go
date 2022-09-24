package model

import (
	"fliteklub/misc"
	"time"
)

type Reservation struct {
	misc.GormModel
	Aircraft Aircraft  `json:"aircraft"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}

func ExampleReservation() Reservation {
	return Reservation{
		GormModel: misc.GormModel{},
		Aircraft:  ExampleAircraft(),
		Start:     misc.OneHourLater(),
		End:       misc.OneHourLater(),
	}

}
