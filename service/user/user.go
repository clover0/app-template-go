package user

import (
	"auth465/core"
	"errors"
	"github.com/jmoiron/sqlx"
)

type userService struct {
	db        *sqlx.DB
	userStore core.UserStore
}

func New(db *sqlx.DB, userStore core.UserStore) core.UserService {
	return userService{
		db:        db,
		userStore: userStore,
	}
}

func (service userService) Register(user *core.User) error {
	//service.userStore.
	tx,_ := service.db.Beginx()
	findFn := service.userStore.Find(1)
	u, _ :=findFn(tx)
	print(u)
	
	tx.Commit()
	return errors.New("TODO!")
}
