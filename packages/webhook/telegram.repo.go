package webhook

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type TelegramRepository interface {
	CreateUser(TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) // Status | url
}
