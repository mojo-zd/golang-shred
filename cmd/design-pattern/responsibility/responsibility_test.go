package responsibility

import "testing"

func TestResponsibility(t *testing.T) {
	manager := &Manager{}
	vice := &ViceManager{Next: manager}
	director := &Director{vice}
	director.HandlerRequest(LeaveRequest{Name: "Mojo", Day: 4, Reason: "有事"})
}
