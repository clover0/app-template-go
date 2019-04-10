package main

import (
	"auth465/config"
	"auth465/db"
	"auth465/store"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var storeSet = wire.NewSet(
	provideDatabase,
	users.New,
)

func provideDatabase(config config.Config) (*sqlx.DB, error) {
	return db.NewDB(config), nil
}
