package session

import (
	"auth465/core"
	"auth465/db"
	"github.com/jmoiron/sqlx"
)

type sessionService struct {
	db            *sqlx.DB
	userStoreFunc core.UserStoreFunc
}

func New(db *sqlx.DB, userStoreFunc core.UserStoreFunc) core.SessionService {
	return sessionService{
		db:            db,
		userStoreFunc: userStoreFunc,
	}
}

func (service sessionService) FindUserByEmail(email string) (*core.User, error) {
	var user *core.User
	err := db.Transact(service.db, func(tx *sqlx.Tx) error {
		var err error
		userStore := service.userStoreFunc(tx)
		user, err = userStore.FindByEmail(email)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

func (service sessionService) FindUserById(id uint32) (*core.User, error) {
	var user *core.User
	err := db.Transact(service.db, func(tx *sqlx.Tx) error {
		var err error
		userStore := service.userStoreFunc(tx)
		user, err = userStore.Find(id)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}