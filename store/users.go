package users

import (
	"auth465/core"
	"github.com/jmoiron/sqlx"
)

func New(tx *sqlx.Tx) *core.UserStore {
	return &userStore{tx}
}

type userStore struct {
	tx *sqlx.Tx
}

func (u *userStore) Find(id int64) (*core.User, error) {
	row := u.tx.QueryRow(FindUser, id)
	var user core.User
	err := row.Scan(&user)
	return &user, err
}

func (u *userStore) Create(user *core.User) error {
	params := toParams(user)
	_, err := u.tx.NamedExec(AddUser, params)
	return err
}

func toParams(user *core.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"updated_at": user.UpdatedAt,
		"created_at": user.CreatedAt,
		"email":      user.Email,
		"password":   user.Password,
	}
}

const FindUser = "SELECT * FROM users WHERE id = ?"

const AddUser = `
INSERT INTO users (
id
,updated_at
,created_at
,email
,password
) VALUES (
:id
,:updated_at
,:created_at
,:email
,:password
)
`
