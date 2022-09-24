package handlers

import (
	"testing"
)

func TestInheritance(t *testing.T) {
	var _ CrudHandler = new(UserHandler)
	var _ CrudHandler = new(ClubHandler)
	var _ CrudHandler = new(AircraftHandler)
}
