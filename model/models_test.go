package model

import (
	"testing"
)

func TestInheritance(t *testing.T) {
	var _ Model = new(User)
	var _ Model = new(Club)
	var _ Model = new(Membership)
	var _ Model = new(Ownership)
	var _ Model = new(Reservation)
}
