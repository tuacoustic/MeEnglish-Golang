package webhook

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type TelegramVieRepository interface {
	// Vie
	StudyNowVie(TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) // Status | url
}
