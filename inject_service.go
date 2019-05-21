package main

import (
	"auth465/core"
	"auth465/service/session"
	"auth465/service/user"

	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var serviceSet = wire.NewSet(
	provideUserService,
	provideSessionService,
)

func provideUserService(db *sqlx.DB, userStoreFunc core.UserStoreFunc) core.UserService {
	return user.New(db, userStoreFunc)
}

func provideSessionService(db *sqlx.DB, s *redis.Client, userStoreFunc core.UserStoreFunc) core.SessionService {
	return session.New(db, s, userStoreFunc)
}

