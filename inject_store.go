package main

import (
	"auth465/config"
	"auth465/core"
	"auth465/db"
	"auth465/store"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var storeSet = wire.NewSet(
	provideDatabase,
	provideUserStore,
)

func provideDatabase(config config.Config) (*sqlx.DB, error) {
	return db.NewDB(config), nil
}

func provideUserStore(db *sqlx.DB) core.UserStore {
	return users.New(db)
}
