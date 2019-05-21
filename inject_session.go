package main

import (
	"auth465/config"
	"auth465/session"

	"github.com/go-redis/redis"
	"github.com/google/wire"
)

var sessionSet = wire.NewSet(
	provideSession,
)

func provideSession(config config.Config) (*redis.Client, error) {
	return session.NewSessionStore(config), nil
}
