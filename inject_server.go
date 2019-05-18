package main

import (
	"auth465/config"
	"auth465/handler/api"
	"auth465/server"

	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	api.New,
	provideServer,
)

func provideServer(config config.Config, api api.Api) *server.Server {
	return &server.Server{
		Handler: "test",
		Api:     api,
	}
}
