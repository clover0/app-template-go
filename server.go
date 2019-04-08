package main

import (
	"auth465/config"
	"auth465/handler/api"
	"auth465/router"

	"log"

	"github.com/labstack/echo"
)

func main() {
	conf := config.InitConfig()
	_, err := InitializeApplication(conf)
	if err != nil {
		log.Fatal("main: cannot initialize application")
	}
	e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	router.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func newApplication(server *Server) application {
	return application{
		server: server,
	}
}

type Server struct {
	Handler string // todo: fix
	Api     api.Api
}

type application struct {
	server *Server
}
