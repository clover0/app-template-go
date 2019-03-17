package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func CreateUserHandler(c echo.Context)(err error) {
	
	return c.JSON(http.StatusOK, 1)
}

func UpdateUserHandler(c echo.Context)(err error){

	return c.JSON(http.StatusOK, 1)
}

