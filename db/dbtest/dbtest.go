package dbtest

import "github.com/jmoiron/sqlx"

// for test

var con *sqlx.DB

func SetCon(c *sqlx.DB) {
	con = c
}

func GetCon() *sqlx.DB {
	if con == nil {
		panic("not exist database connection")
	}
	return con
}
