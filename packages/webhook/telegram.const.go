package webhook

import (
	"me-english/utils/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type textHandlingStruct struct {
	StartBot       string
	StudyNowVie    string
	AutoRemindVie  string
	InstructionVie string
	SupportVie     string
	DevelopVie     string
	DonateVie      string
}

var (
	msg            = ""
	telegramParams = config.SendTelegramMsgStruct{
		ChatID:      uint64(664743441),
		Text:        "Lỗi",
		ReplyMarkup: "",
		ParseMode:   "markdown",
	}
	Command_Handling = textHandlingStruct{
		StartBot:       "/",
		StudyNowVie:    "học ngay",
		AutoRemindVie:  "nhắc học tự động",
		InstructionVie: "hướng dẫn",
		SupportVie:     "gửi hỗ trợ",
		DevelopVie:     "cùng phát triển",
		DonateVie:      "ủng hộ tác giả",
	}
	Home_Reply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Học ngay"),
			tgbotapi.NewKeyboardButton("Nhắc học tự động"),
			tgbotapi.NewKeyboardButton("Hướng dẫn"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Gửi hỗ trợ"),
			tgbotapi.NewKeyboardButton("Cùng phát triển"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Ủng hộ tác giả"),
		),
	)
)
