package api

import "auth465/core"

func New(
	userStoreFunc core.UserStoreFunc,
	userService core.UserService,
) Api {
	return Api{
		UserStoreFunc: userStoreFunc,
		UserService: userService,
	}
}

type Api struct {
	UserStoreFunc core.UserStoreFunc
	UserService core.UserService
}
