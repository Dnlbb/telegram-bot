package telegram

import "github.com/Dnlbb/telegram-bot/internal/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// stor
}

// func New(client *telegram.Client, stor storage )
