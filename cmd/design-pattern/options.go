package design_pattern

import (
	"github.com/rs/zerolog/log"
)

type Options struct {
	timeout int
	factory string
}

type Option func(options *Options)

func newService(opt ...Option) {
	opts := newOptions(opt...)
	log.Info().
		Interface("timeout", opts.timeout).
		Str("factory", opts.factory).
		Msg("display opts info")
}

func newOptions(opt ...Option) Options {
	opts := Options{
		timeout: 5,
		factory: "bridge",
	}
	for _, o := range opt {
		o(&opts)
	}
	return opts
}
