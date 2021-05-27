package webhook

import (
	"fmt"
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
		InstructionVie: "từ vựng đã lưu",
		SupportVie:     "gửi hỗ trợ",
		DevelopVie:     "cùng phát triển",
		DonateVie:      "ủng hộ tác giả",
	}
	Home_Reply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Học ngay"),
			tgbotapi.NewKeyboardButton("Nhắc học tự động"),
			tgbotapi.NewKeyboardButton("Từ vựng đã lưu"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Gửi hỗ trợ"),
			tgbotapi.NewKeyboardButton("Cùng phát triển"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Ủng hộ tác giả"),
		),
	)
	StudyNowVie_Reply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Học Group 1"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("﹒1﹒"),
			tgbotapi.NewKeyboardButton("2 >"),
			tgbotapi.NewKeyboardButton("3 >"),
			tgbotapi.NewKeyboardButton("4 >"),
			tgbotapi.NewKeyboardButton("10 >>"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("﹒Gr1﹒"),
			tgbotapi.NewKeyboardButton("Gr2 >"),
			tgbotapi.NewKeyboardButton("Gr3 >"),
			tgbotapi.NewKeyboardButton("Gr10 >>"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back Home"),
		),
	)
)

func StudyNowVie(AwlGroupID uint64, groupVocab string) string {
	return fmt.Sprintf(`
*Group %d*

Lưu ý: Các bạn có thể click bên cạnh từ vựng để thấy chi tiết từ vựng đó nhé. Mỗi trang sẽ chứa 15 từ vựng theo Group, chúc các bạn học tập hiệu quả 😉

%s

*Tip theo nút*:
1. Học theo Group đang xem
2. Trang từ vựng theo Group
3. Trang Group
4. Trở về Trang chính
`, AwlGroupID, groupVocab)
}
