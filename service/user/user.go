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

// CheckDuplicateEmail returns true if not duplicate email
func (service userService) CheckDuplicateEmail(user *core.User) (bool, error) {
	var res bool
	err := db.Transact(service.db, func(tx *sqlx.Tx) error {
		var err error
		res = true

		userStore := service.userStoreFunc(tx)

		count, err := userStore.Count("email", user.Email)
		if err != nil {
			log.Error("cannot count user by email")
			res = false
			return err
		}
		if count > 0 {
			res = false
		}

		return nil
	})
	return res, err
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
