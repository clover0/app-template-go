package main

import (
	"auth465/config"
	"auth465/router"
	"auth465/server"

	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	conf := config.InitConfig("")
	app, err := InitializeApplication(conf)
	if err != nil {
		log.Fatal("main: cannot initialize application")
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: include initApp
	router.Init(e, app.server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newApplication(server *server.Server) application {
	return application{
		server: server,
	}
}

type application struct {
	server *server.Server
}
