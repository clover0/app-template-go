package main

import (
	"auth465/core"
	"auth465/service/user"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var serviceSet = wire.NewSet(
	provideUserService,
)

func provideUserService(db *sqlx.DB, userStoreFunc core.UserStoreFunc) core.UserService {
	return user.New(db, userStoreFunc)
}

