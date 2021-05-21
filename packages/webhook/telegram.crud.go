package webhook

import (
	"errors"
	"me-english/entities"
	"me-english/utils/channels"
	"me-english/utils/config"
	"me-english/utils/console"
	sendReq "me-english/utils/sendRequest"
	"net/url"

	"github.com/jinzhu/gorm"
)

type repositoryTelegramCRUD struct {
	db *gorm.DB
}

func NewRepositoryTelegramCRUD(db *gorm.DB) *repositoryTelegramCRUD {
	return &repositoryTelegramCRUD{db}
}

var (
	msg = ""
)

func (r *repositoryTelegramCRUD) CreateUser(userData TelegramRespJSON) (bool, error) {
	// var err error
	// countExistTelegramUser := entities.CountTelegramUsers{}
	countCurrentUsers := entities.CountTelegramUsers{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// createData := entities.TelegramUsers{
		// 	TelegramID: userData.Message.From.ID,
		// 	FirstName:  userData.Message.From.FirstName,
		// 	LastName:   userData.Message.From.LastName,
		// 	Username:   userData.Message.From.UserName,
		// 	Type:       userData.Message.Chat.Type,
		// }
		// queryDBByTelegramID := QueryExistTelegramID(userData.Message.From.ID)
		// r.db.Raw(queryDBByTelegramID).Find(&countExistTelegramUser)
		// if countExistTelegramUser.Count > 0 {
		// 	msg = "Tên Telegram này đã tồn tại hệ thống"
		// 	ch <- false
		// 	return
		// }
		// err = r.db.Debug().Model(&entities.TelegramUsers{}).Create(&createData).Error
		// if err != nil {
		// 	msg = fmt.Sprintf("%s", err)
		// 	ch <- false
		// 	return
		// }
		queryDBForCountAllUsers := QueryAllTelegramUsers()
		r.db.Raw(queryDBForCountAllUsers).Find(&countCurrentUsers)
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		// Gửi tin Telegram thông báo
		telegramMsg := url.QueryEscape(ParamTelegramSendTextWelcome(userData.Message.From.UserName, int64(countCurrentUsers.Count)))
		replyMarkup := url.QueryEscape(ParamTelegramSendReplyMarkupWelcome())
		telegramParams := config.SendTelegramMsgStruct{
			ChatID:      userData.Message.From.ID,
			Text:        telegramMsg,
			ReplyMarkup: replyMarkup,
			ParseMode:   "markup",
		}
		getTelegramMsgUrl := config.GetTelegramMeEnglishSendMsgUrlConfig(telegramParams)
		responseData := sendReq.PostRequestToTelegram(getTelegramMsgUrl, "GET", "")
		console.Info(responseData)
		console.Info(telegramMsg)
		return true, nil
	}
	return false, errors.New(msg)
}
