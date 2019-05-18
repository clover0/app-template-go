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
		Count(column string, param interface{}) (uint, error)
		Find(uint32) (*User, error)
		Create(user *User) (uint32, error)
	}

	// improve
	UserStoreFunc func(session *sqlx.Tx) UserStore

	UserService interface {
		CheckDuplicateEmail(user *User) (bool,error)
		Register(user *User) error
	}
)
