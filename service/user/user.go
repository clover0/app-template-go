package user

import (
	"auth465/core"
	"errors"
	"github.com/jmoiron/sqlx"
)

type userService struct {
	db            *sqlx.DB
	userStoreFunc core.UserStoreFunc
}

func New(db *sqlx.DB, userStoreFunc core.UserStoreFunc) core.UserService {
	return userService{
		db:            db,
		userStoreFunc: userStoreFunc,
	}
}

func (service userService) Register(user *core.User) error {
	//service.userStore.
	tx, _ := service.db.Beginx()
	userStore := service.userStoreFunc(tx)
	user, err := userStore.Find(1)
	print(err)

	tx.Commit()
	return errors.New("TODO!")
}
