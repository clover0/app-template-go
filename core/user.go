package core

import (
	"github.com/jmoiron/sqlx"
)

type (
	User struct {
		ID        uint32
		UpdatedAt string
		CreatedAt string
		Email     string
		Password  string
	}

	UserStore interface {
		Find(uint32) (*User, error)
		Create(user *User) (uint32, error)
	}

	// improve
	UserStoreFunc func(session *sqlx.Tx) UserStore

	UserService interface {
		Register(user *User) error
	}
)
