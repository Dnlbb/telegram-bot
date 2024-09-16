package config

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var (
	Redis struct {
		Password string
		DB       int
	}
	Token string
	Host  string
	Stor  string
)

func Init() {
	Redis.Password = os.Getenv("PASSWORD")
	if dbStr := os.Getenv("DB"); dbStr != "" {
		if db, err := strconv.Atoi(dbStr); err == nil {
			Redis.DB = db
		} else {
			log.Fatalf("Invalid DB value: %v", err)
		}
	}

	flag.StringVar(&Token, "token-bot-token", "", "token for access telegram bot")
	flag.StringVar(&Host, "host-bot", "", "host for telegram bot")
	flag.StringVar(&Stor, "storage", "", "need storage")
	flag.Parse()

	if Stor == "" {
		log.Fatal("Need storage")
	}

	if Token == "" || Host == "" {
		log.Fatal("Need token and host")
	}
}
