package api

import "auth465/core"

func New(
	userStore core.UserStore,
) Api {
	return Api{
		UserStore: userStore,
	}
}

type Api struct {
	UserStore core.UserStore
}
