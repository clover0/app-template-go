package config

import (
	"log"
	"os"
	"github.com/BurntSushi/toml"
)

// from toml file config.
type Config struct {
	DB      DataBaseConfig     `toml:"database"`
	Session SessionStoreConfig `toml:"session"`
}

type DataBaseConfig struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	DbName      string `toml:"dbName"`
	MaxIdleConn int    `toml:"maxIdleConn"`
	MaxOpenConn int    `toml:"maxOpenConn"`
}

type SessionStoreConfig struct {
	Addr string `toml:"addr"`
	DB   int    `toml:"db"`
}

// Initialization configuration for project 
// every environment.
func InitConfig() Config {
	var config Config
	var configFile string
	configDir := "config"

	switch os.Getenv("AUTH465_ENV") {
	case "local":
		configFile = "config.local.toml"
	case "development":
		configFile = "config.development.toml"
	case "staging":
		configFile = "config.staging.toml"
	case "production":
		configFile = "config.production.toml"
	default:
		configFile = "config.local.toml"
	}

	_, err := toml.DecodeFile(configDir+"/"+configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	if config.DB.Password == "" {
		os.Getenv("AUTH465_DB_PWD")
	}

	return config
}
