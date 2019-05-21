package session

import (
	"auth465/core"
	"auth465/db"

	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/go-redis/redis"
	"github.com/labstack/gommon/log"
	"github.com/jmoiron/sqlx"
)

const defaultSessionExpiration = 1 * time.Hour

type sessionService struct {
	db            *sqlx.DB
	sessionStore  *redis.Client
	userStoreFunc core.UserStoreFunc
}

func New(db *sqlx.DB, sessionStore *redis.Client, userStoreFunc core.UserStoreFunc) core.SessionService {
	return sessionService{
		db:            db,
		sessionStore:  sessionStore,
		userStoreFunc: userStoreFunc,
	}
}

// CreateSession saves session to session-store
func (service sessionService) CreateSession(userId uint32) (string, error) {
	sid := generateSessionId()
	err := service.sessionStore.Set(sid, userId, defaultSessionExpiration).Err()

	return sid, err
}

// FindUserByEmail finds user by email and return user with transaction
func (service sessionService) FindUserByEmail(email string) (*core.User, error) {
	var user *core.User
	err := db.Transact(service.db, func(tx *sqlx.Tx) error {
		var err error
		userStore := service.userStoreFunc(tx)
		user, err = userStore.FindByEmail(email)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

// FindUserById finds user by id and return user with transaction
func (service sessionService) FindUserById(id uint32) (*core.User, error) {
	var user *core.User
	err := db.Transact(service.db, func(tx *sqlx.Tx) error {
		var err error
		userStore := service.userStoreFunc(tx)
		user, err = userStore.Find(id)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

func (service sessionService) ComparePassword(user *core.User, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input))
}

func generateSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Error("can not generate sessionId")
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
