package config

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

type ConfigList struct {
	DbUser    string
	Password  string
	Tcp       string
	DbName    string
	SQLDriver string

	Port int

	APIKey string
}

var Config ConfigList

func Init() {
	cfg, err := ini.Load("secret.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		DbUser:    cfg.Section("db").Key("dbuser").String(),
		Password:  cfg.Section("db").Key("password").String(),
		Tcp:       cfg.Section("db").Key("tcp").String(),
		DbName:    cfg.Section("db").Key("dbname").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Port:      cfg.Section("web").Key("port").MustInt(),
		APIKey:    cfg.Section("sendapi").Key("apikey").String(),
	}
}
