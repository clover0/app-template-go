package users

import (
	"auth465/core"

	"crypto/rand"
	"encoding/binary"

	"github.com/jmoiron/sqlx"
)

func New() core.UserStoreFunc {
	return func(session *sqlx.Tx) core.UserStore {
		return &userStore{
			sess: session,
		}
	}
}

type userStore struct {
	sess *sqlx.Tx
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

func (u *userStore) Create(user *core.User) (uint64, error) {
	var err error
	user.ID = generateUserId()
	params := toCreateParams(user)
	_, err = u.sess.NamedExec(AddUser, params)
	if err != nil {
		u.sess.Rollback()
		return 0, err
	}

	return user.ID, u.sess.Commit()
}

func toCreateParams(user *core.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"updated_at": user.UpdatedAt,
		"created_at": user.CreatedAt,
		"email":      user.Email,
		"password":   user.Password,
	}
}

func generateUserId() uint64 {
	bs := make([]byte, 128)
	if _, err := rand.Read(bs); err != nil {
		panic("")
	}
	return binary.BigEndian.Uint64(bs)
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
