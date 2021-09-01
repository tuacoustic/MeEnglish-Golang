package webhook

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type TelegramVieRepository interface {
	// Vie
	GetStudyNowVie(TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup)
	GetVocabByGroupPageVie(TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup)
	GetVocabByGroupVie(TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup)
	BackHomePage(TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup)
	GroupStudy(TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup)
	FindVocab(TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup)
	FindAudio(TelegramRespJSON) (bool, string, error)
	FindImage(TelegramRespJSON) (bool, string, error)
	AnswerQuestionButton(TelegramRespJSON) (bool, string, error)
	HandleTrueAnswer(TelegramRespJSON) (bool, string, error, tgbotapi.ReplyKeyboardMarkup)
	AnswerQuestionByText(TelegramRespJSON) (bool, string, error)
	ShowAnswer(TelegramRespJSON) (bool, string, error)
}
