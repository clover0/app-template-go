package db

import (
	"fmt"
	"log"

	"auth465/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbcon *sqlx.DB

func NewDB(config2 config.Config) *sqlx.DB {
	user := config2.DB.User
	host := config2.DB.Host
	port := config2.DB.Port
	password := config2.DB.Password
	dbName := config2.DB.DbName
	maxIdleConn := config2.DB.MaxIdleConn
	maxOpenConn := config2.DB.MaxOpenConn

	ds := fmt.Sprintf(
		"user=%s host=%s port=%d password=%s dbname=%s sslmode=disable",
		user,
		host,
		port,
		password,
		dbName,
	)
	db, err := sqlx.Connect("postgres", ds)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)
	log.Printf("database connected: user=%s host=%s port=%d password=*** dbname=%s sslmode=disable maxidle=%d maxopen=%d",
		user,
		host,
		port,
		dbName,
		maxIdleConn,
		maxOpenConn,
	)

	dbcon = db

	return db
}

func GetConnection() *sqlx.DB {
	return dbcon
}