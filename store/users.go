package users

import (
	"auth465/core"

	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) core.UserStore {
	return &userStore{db}
}

type userStore struct {
	db *sqlx.DB
}

func (u *userStore) Find(id int64) (*core.User, error) {
	var err error
	tx, err := u.db.Beginx()
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow(FindUser, id)
	var user core.User
	err = row.Scan(&user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, tx.Commit()
}

func (u *userStore) Create(user *core.User) error {
	var err error
	params := toParams(user)
	tx, err := u.db.Beginx()

	_, err = tx.NamedExec(AddUser, params)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
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
