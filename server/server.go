package server

import (
	"auth465/handler/api"

	"github.com/go-redis/redis"
)

type Server struct {
	Handler string // todo: fix
	Api     *api.Api
	SessionStore *redis.Client
}