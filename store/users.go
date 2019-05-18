package users

import (
	"auth465/core"

	"time"
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

func (u *userStore) Find(id uint32) (*core.User, error) {
	var err error
	row := u.sess.QueryRow(FindUser, id)
	var user core.User
	err = row.Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userStore) Create(user *core.User) (uint32, error) {
	var err error
	user.ID = generateUserId()
	params := toCreateParams(user)
	_, err = u.sess.NamedExec(AddUser, params)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func toCreateParams(user *core.User) map[string]interface{} {
	now := time.Now()
	timestamp := now.Format(time.RFC3339Nano)
	user.UpdatedAt = timestamp
	user.CreatedAt = timestamp
	return map[string]interface{}{
		"id":         user.ID,
		"updated_at": user.UpdatedAt,
		"created_at": user.CreatedAt,
		"email":      user.Email,
		"password":   user.Password,
	}
}

func generateUserId() uint32 {
	bs := make([]byte, 128)
	if _, err := rand.Read(bs); err != nil {
		panic("")
	}
	return binary.BigEndian.Uint32(bs)
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
