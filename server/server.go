package server

import "auth465/handler/api"

type Server struct {
	Handler string // todo: fix
	Api     api.Api
}