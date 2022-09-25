package model

import (
	"testing"
)

func TestInheritance(t *testing.T) {
	var _ Model = new(Aircraft)
	var _ Model = new(Airport)
	var _ Model = new(Club)
	var _ Model = new(Reservation)
	var _ Model = new(User)
}
