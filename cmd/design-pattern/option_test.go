package design_pattern

import (
	"testing"
)

func TestOptions(t *testing.T) {
	opts := []Option{func(options *Options) {
		options.timeout = 10
	}, func(options *Options) {
		options.factory = "host"
	}}
	newService(opts...)
}
