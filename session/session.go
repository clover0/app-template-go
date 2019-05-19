package session

import (
	"auth465/config"
	"log"

	"github.com/go-redis/redis"
)

func NewSessionStore(config config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Session.Addr,
		Password: "", // no password set
		DB:       config.Session.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	log.Printf("session connected: Addr=%s db=%d",
		config.Session.Addr,
		config.Session.DB,
	)

	return client
}
