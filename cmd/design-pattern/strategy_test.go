package design_pattern

import "testing"

func TestStrategy(t *testing.T)  {
	NewPaymentContext("mojo", "123456", 100, &Cash{}).Pay()
	NewPaymentContext("kimo", "234567", 99, &Bank{}).Pay()
}