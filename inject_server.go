package main

import (
	"auth465/config"
	"auth465/handler/api"

	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	api.New,
	provideServer,
)

func provideServer(config config.Config, api api.Api) *Server {
	return &Server{
		Handler: "test",
		Api:     api,
	}
}
