package state

import "testing"

func TestState(t *testing.T) {
	ctx := NewHotelContext()
	ctx.Handler(Free)
	ctx.Handler(Booked)
	ctx.Handler(Checked)
}
