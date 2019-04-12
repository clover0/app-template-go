package users

import (
	"auth465/core"
	"auth465/db"
)

func New() core.UserStoreFunc {
	return func(session core.StoreSession) core.UserStore {
		return &userStore{
			sess: session.(db.Session),
		}
	}
}

type userStore struct {
	sess db.Session
}

func (u *userStore) Find(id int64) (*core.User, error) {
	var err error
	row := u.sess.QueryRow(FindUser, id)
	var user core.User
	err = row.Scan(&user)
	if err != nil {
		u.sess.Rollback()
		return nil, err
	}
	return &user, nil
}

//func (u *userStore) Create(user *core.User) error {
//	var err error
//	params := toParams(user)
//	tx, err := u.db.Beginx()
//
//	_, err = tx.NamedExec(AddUser, params)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	return tx.Commit()
//}

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
