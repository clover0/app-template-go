package handler

import (
	"auth465/core"
	"fmt"

	"strconv"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
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
				cookie.Path = "/"
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

func DeleteSessionHandler(session *redis.Client) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		cookie, err := c.Cookie("ESESSION")
		if err == http.ErrNoCookie {
			return c.JSON(http.StatusOK, "no cookie")
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, "can not read cookie")
		}
		session.Del(cookie.Value)
		return c.JSON(http.StatusOK, "ok")
	}
}

func ShowCurrentSessionHandler(service core.SessionService, session *redis.Client) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		cookie, err := c.Cookie("ESESSION")
		if err == http.ErrNoCookie {
			return c.JSON(http.StatusOK, "no sign in")
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, "can not read cookie")
		}
		userId, err := session.Get(cookie.Value).Result()
		if err == redis.Nil {
			return c.JSON(http.StatusOK, "no sign in")
		} else if err != nil {
			log.Error(err)
			panic(err)
		}
		log.Info(userId)
		id, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			log.Error(err)
			panic(err)
		}
		user, err := service.FindUserById(uint32(id))
		log.Info(id)
		log.Info(uint32(id))
		if user == nil {
			return c.JSON(http.StatusInternalServerError, "error")
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("already sign in! Your email is %s", user.Email))
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
