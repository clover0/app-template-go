package handler

import (
	"auth465/core"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"

	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"
)

type SessionCreateForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateSessionHandler(service core.SessionService, session *redis.Client) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {

		// bind form from request
		form := new(UserCreateForm)
		if err := c.Bind(form); err != nil {
			return err
		}

		// check database
		user, err := service.FindUserByEmail(form.Email)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, "error")
		}
		// exists user?
		if user != nil {
			if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)) == nil {
				sessionId := generateSessionId()
				session.Set(sessionId, user.ID, 1*time.Hour)
				cookie := new(http.Cookie)
				cookie.Name = "ESESSION"
				cookie.Value = sessionId
				cookie.Expires = time.Now().Add(14 * 24 * time.Hour)
				c.SetCookie(cookie)
				return c.JSON(http.StatusCreated, 1)
			} else { // exists user but password no match
				return c.JSON(http.StatusBadRequest, "password no match")
			}
		} else { // not exists
			return c.JSON(http.StatusBadRequest, "email no match")
		}
	}
}

func generateSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Error("can not generate sessionId")
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
