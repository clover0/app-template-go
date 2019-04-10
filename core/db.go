package core

import "github.com/jmoiron/sqlx"

type Session interface {
	WithSession(db *sqlx.DB) error
}