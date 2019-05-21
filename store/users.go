package users

import (
	"auth465/core"
	"crypto/rand"
	"encoding/binary"
	"time"

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

// Find finds user from users table by id(PK)
func (u *userStore) Find(id uint32) (*core.User, error) {
	var user core.User
	err := u.sess.QueryRowx(FindUser, id).StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail finds user from users table by email(UK)
func (u *userStore) FindByEmail(email string) (*core.User, error) {
	var user core.User
	err := u.sess.QueryRowx(FindUserByEmail, email).StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Count counts from users table by 'where column = param'
func (u *userStore) Count(column string, param interface{}) (uint, error) {
	var err error
	var res uint
	stmt := countByColumnNameStmt(column)
	err = u.sess.Get(&res, stmt, param)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// Create inserts to users table by parameter user
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

// toCreateParams creates map for insert statement
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

// generateUserId generates random number for user-id
func generateUserId() uint32 {
	bs := make([]byte, 256)
	if _, err := rand.Read(bs); err != nil {
		panic("")
	}

	return binary.BigEndian.Uint32(bs)
}

// count statement
func countByColumnNameStmt(column string) string {
	return "SELECT count(*) FROM users WHERE " + column + " = $1"
}

const FindUser = "SELECT * FROM users WHERE id = $1"
const FindUserByEmail = "SELECT * FROM users WHERE email = $1"

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
