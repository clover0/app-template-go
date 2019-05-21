package api

import "auth465/core"

func New(
	userStoreFunc core.UserStoreFunc,
	userService core.UserService,
	sessionService core.SessionService,
) *Api {
	return &Api{
		UserStoreFunc: userStoreFunc,
		UserService: userService,
		SessionService: sessionService,
	}
}

type Api struct {
	UserStoreFunc core.UserStoreFunc
	UserService core.UserService
	SessionService core.SessionService
}
