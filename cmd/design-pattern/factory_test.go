package design_pattern

import "testing"

func TestFactory(t *testing.T)  {
	NewAPI(func(options *Options) {
		options.factory = "bridge"
	}).SayHi("test")

	NewAPI(func(options *Options) {
		options.factory = "host"
	}).SayHi("test")
}