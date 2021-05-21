package config

import (
	"fmt"
	"strings"
)

type SendTelegramMsgStruct struct {
	ChatID      uint64
	Text        string
	ReplyMarkup string
	ParseMode   string
}

func GetTelegramMeEnglishSendMsgUrlConfig(sendMsg SendTelegramMsgStruct) string {
	inputToken := strings.Replace(TELEGRAM_SEND_MESSAGE, "@token_params", TELEGRAM_TOKEN_MEENGLISH, -1)
	url := strings.Replace(inputToken, "@params", fmt.Sprintf("chat_id=%d&text=%s&reply_markup=%s&parse_mode=%s", sendMsg.ChatID, sendMsg.Text, sendMsg.ReplyMarkup, sendMsg.ParseMode), -1)
	return url
}
