package handlers

import (
	"testing"
)

func TestInheritance(t *testing.T) {
	var _ CrudHandler = new(AircraftHandler)
	var _ CrudHandler = new(ClubHandler)
	var _ CrudHandler = new(ReservationHandler)
	var _ CrudHandler = new(UserHandler)
}
