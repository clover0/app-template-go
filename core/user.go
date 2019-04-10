package core

import "github.com/jmoiron/sqlx"

type (
	User struct {
		ID        int64
		UpdatedAt string
		CreatedAt string
		Email     string
		Password  string
	}
	UserStore interface {
		Find(int64) func (*sqlx.Tx) (*User, error)
		//Create(user *User) error
	}

	UserService interface {
		Register(user *User) error
	}
)
