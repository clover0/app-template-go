package user

import (
	"auth465/core"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
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
	tx, _ := service.db.Beginx()
	userStore := service.userStoreFunc(tx)
	userId, err := userStore.Create(user)
	if err != nil {
		log.Error("cannot create user")
		return err
	}
	log.Info("create user", userId)

	tx.Commit()
	return errors.New("TODO!")
}
