package design_pattern

import "github.com/rs/zerolog/log"

type API interface {
	SayHi(content string)
}

type hostAPI struct {
}

func (*hostAPI) SayHi(content string) {
	log.Info().Msg("call host api SayHi method!")
}

func NewAPI(options ...Option) API {
	opts := &Options{
		timeout:5,
		factory:"host",
	}
	for _, o := range options {
		o(opts)
	}
	var api API
	switch opts.factory {
	case "host":
		api = &hostAPI{}
	case "bridge":
		api = &bridgeAPI{}
	}
	return api
}

type bridgeAPI struct {
}

func (*bridgeAPI) SayHi(content string)  {
	log.Info().Msg("call bridge api SayHi method!")
}