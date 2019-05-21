package server

import (
	"auth465/server/api"

	"github.com/go-redis/redis"
)

type Server struct {
	Handler string // todo: fix
	Api     *api.Api
	SessionStore *redis.Client
}