package router

import (
	"auth465/handler"
	"auth465/server"

	"net/http"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo, s *server.Server)  {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/a/user/create",
		handler.CreateUserHandler(s.Api.UserService, s.SessionStore),
	)
	//e.POST("/a/u/u", handler.CreateUserHandler)
	//e.POST("/a/s/c", handler.CreateUserHandler)
	//e.POST("/a/s/d", handler.CreateUserHandler)
}