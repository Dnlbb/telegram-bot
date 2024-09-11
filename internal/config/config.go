package config

import (
	"flag"
	"log"
)

func Token() string {
	token := flag.String("token-bot-token", "", "token for access telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("Need token")
	}
	return *token
}

func Host() string {
	host := flag.String("host-bot", "", "host for telegram bot")
	flag.Parse()
	if *host == "" {
		log.Fatal("Need host")
	}
	return *host
}
