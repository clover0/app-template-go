package main

import (
	"auth465/router"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	router.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}
