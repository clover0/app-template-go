package main

import (
	"auth465/config"
	"auth465/db"
	"auth465/router"
	"github.com/labstack/echo"
)

func main() {
	initializeApp()
	e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	router.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func initializeApp()  {
	config.InitConfig()
	db.NewDB()
}
