package core

import (
	"github.com/jmoiron/sqlx"
)

type (
	User struct {
		ID        uint64
		UpdatedAt string
		CreatedAt string
		Email     string
		Password  string
	}
	UserStore interface {
		Find(int64) (*User, error)
		Create(user *User) (uint64, error)
	}

	// improve
	UserStoreFunc func(session *sqlx.Tx) UserStore

	UserService interface {
		Register(user *User) error
	}
)
