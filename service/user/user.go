package user

import (
	"auth465/core"
	"auth465/db"
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
	return db.Transact(service.db, func(tx *sqlx.Tx) error {
		userStore := service.userStoreFunc(tx)
		
		userId, err := userStore.Create(user)
		if err != nil {
			log.Error("cannot create user")
			return err
		}
		log.Info("registered user: ", userId)
		return nil
	})
}
