package webhook

import "fmt"

var (
	Author = "*From TuDinh*"
)

func ParamTelegramSendTextWelcome(telegramName string, userNumber int64) string {
	return fmt.Sprintf(`
%s

Yay *%s*!
Chúc mừng bạn là người thứ: *%d* kích hoạt BOT học tập này
🎉 Mình xin cảm ơn đội ngũ đã giúp mình hoàn thành BOT này:
❤️ Từ điển, Oxford: https://www.oxfordlearnersdictionaries.com
❤️ Phát âm Mỹ, Oxford: https://www.oxfordlearnersdictionaries.com
❤️ Hình ảnh chú thích, Pexels: https://pexels.com 
❤️ Thao tác BOT, Telegram: https://telegram.org
Và gia đình, những người bạn, đồng nghiệp, Sếp đã luôn ở bên động viên Tú.
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
			"text": "Từ vựng đã lưu",
			"callback_data": "/instruction"
		}
		],
		[
		{
			"text": "Gửi hỗ trợ",
			"callback_data": "/support"
		},
		{
			"text": "Cùng phát triển",
			"callback_data": "/develop"
		}
		],
		[
			{
				"text": "Ủng hộ tác giả",
				"callback_data": "/donate"
			}
		]
	],
	"resize_keyboard": true,
	"one_time_keyboard": false,
	"selective": true
}
`)
}

func ParamTelegramSendStudyAnswer(telegramName string, userNumber int64) string {
	return fmt.Sprintf(`
*📌 Bạn đang học Group 1*

Vui lòng cung cấp đáp án dưới:
🔑 ----- (##) (##): Khu vực kinh tế

*Definition*
Noun: an area or portion that is distinct from others

*Example*
Noun: operations in the southern ***** of the North Sea

A. Sector
B. Available
C. Financal
D. Process
`)
}
