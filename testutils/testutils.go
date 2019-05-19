package testutils

import (
	"auth465/config"
	"auth465/db"
	"auth465/db/dbtest"
	"crypto/rand"
	"encoding/binary"

	"fmt"
	"os"
)

func SetUpConfig(baseDir string) *config.Config {
	err := os.Setenv("AUTH465_ENV", "test")
	if err != nil {
		fmt.Print(err)
	}
	config := config.InitConfig("../config")
	return &config
}

func SetupDBWithConfig(conf *config.Config) {
	err := os.Setenv("AUTH465_ENV", "test")
	if err != nil {
		fmt.Print(err)
	}
	c := db.NewDB(*conf)
	dbtest.SetCon(c)
}

func CleanupDB() {
	err := dbtest.GetCon().Close()
	fmt.Println("close db connection")
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateRandomNum() uint32 {
	bs := make([]byte, 32)
	if _, err := rand.Read(bs); err != nil {
		panic("")
	}

	return binary.BigEndian.Uint32(bs)
}
