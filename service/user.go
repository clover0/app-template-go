package service

import (
	"auth465/core"
	"errors"
)

type userService struct {
	userStore core.UserStore
}

func (service userService)Register() error {
	//service.userStore.
	return errors.New("TODO!")
}