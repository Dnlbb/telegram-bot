package main

import (
	"log"

	tg "github.com/Dnlbb/telegram-bot/internal/events/telegram"
	"github.com/Dnlbb/telegram-bot/internal/storage"

	"github.com/Dnlbb/telegram-bot/internal/clients/telegram"
	"github.com/Dnlbb/telegram-bot/internal/config"
	event_consumer "github.com/Dnlbb/telegram-bot/internal/consumer/event-consumer"
	fiilestorage "github.com/Dnlbb/telegram-bot/internal/storage/fiileStorage"
	redisstorage "github.com/Dnlbb/telegram-bot/internal/storage/redisStorage"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	redisAddr   = "localhost:6379"
	batchSize   = 100
)

func main() {
	var storage storage.Storage
	config.Init()
	switch config.Stor {
	case "Redis":
		storage = redisstorage.New(redisAddr)
	case "File":
		storage = fiilestorage.New(storagePath)
	}
	eventsProcessor := tg.New(telegram.New(tgBotHost, config.Token), storage)
	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stoppsed", err)
	}

}
