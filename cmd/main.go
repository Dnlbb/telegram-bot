package main

import (
	"log"

	tg "github.com/Dnlbb/telegram-bot/internal/events/telegram"

	"github.com/Dnlbb/telegram-bot/internal/clients/telegram"
	"github.com/Dnlbb/telegram-bot/internal/config"
	event_consumer "github.com/Dnlbb/telegram-bot/internal/consumer/event-consumer"
	fiilestorage "github.com/Dnlbb/telegram-bot/internal/storage/fiileStorage"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := tg.New(telegram.New(tgBotHost, config.Token()), fiilestorage.New(storagePath))
	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stoppsed", err)
	}

}
