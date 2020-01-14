package design_pattern

import "testing"

func TestCommand(t *testing.T)  {
	mb := &MonitorBoard{}
	sc := NewStartCommand(mb)
	rc := NewRebootCommand(mb)
	box := NewBox(sc, rc)

	box.PressButton1()
	box.PressButton2()
}