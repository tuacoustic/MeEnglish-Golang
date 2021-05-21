package webhook

import "fmt"

var (
	Author = "```From TuDinh```"
)

func ParamTelegramSendTextWelcome(telegramName string, userNumber int64) string {
	return fmt.Sprintf(`
%s
Yay *%s*!
Chúc mừng bạn là người thứ: *%d* kích hoạt BOT học tập này
👉  Theo dõi mình nhé
Website : https://tudinh.vn
Facebook: https://fb.com/tudinhacoustic
`, Author, telegramName, userNumber)
}

func ParamTelegramSendReplyMarkupWelcome() string {
	return fmt.Sprintf(`
{
	"keyboard": [
		[
		{
			"text": "Học ngay",
			"callback_data": "/study"
		},
		{
			"text": "Nhắc học tự động",
			"callback_data": "/notification"
		},
		{
			"text": "Cài đặt",
			"callback_data": "/setting"
		}
		],
		[
		{
			"text": "Gửi hỗ trợ",
			"callback_data": "/support"
		},
		{
			"text": "Thoát tác vụ",
			"callback_data": "/quit"
		}
		]
	],
	"resize_keyboard": true,
	"one_time_keyboard": false,
	"selective": true
}
`)
}
