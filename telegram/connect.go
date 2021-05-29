package telegram

import (
	"log"
	"me-english/utils/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ConnectBot() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_TOKEN_MEENGLISH)
	if err != nil {
		return nil, err
	}
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot, nil
}
