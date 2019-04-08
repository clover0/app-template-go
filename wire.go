//+build wireinject

package main

import (
	"auth465/config"

	"github.com/google/wire"
)

func InitializeApplication(config config.Config) (application, error){
	wire.Build(
		//serviceSet,
		storeSet,
		serverSet,
		newApplication,
	)
	return application{}, nil
}