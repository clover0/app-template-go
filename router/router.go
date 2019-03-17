package router

import (
	"auth465/handler"
	"net/http"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo)  {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/a/u/c", handler.CreateUserHandler)
	e.POST("/a/u/u", handler.CreateUserHandler)
	e.POST("/a/s/c", handler.CreateUserHandler)
	e.POST("/a/s/d", handler.CreateUserHandler)
}