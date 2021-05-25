package webhook

import (
	"fmt"
	"me-english/entities"
	"me-english/utils/channels"
	"me-english/utils/config"
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
	msg            = ""
	telegramParams = config.SendTelegramMsgStruct{
		ChatID:      uint64(664743441),
		Text:        "Lỗi",
		ReplyMarkup: "",
		ParseMode:   "markdown",
	}
)

func (r *repositoryTelegramCRUD) CreateUser(userData TelegramRespJSON) (bool, string) {
	var err error
	countExistTelegramUser := entities.CountTelegramUsers{}
	countCurrentUsers := entities.CountTelegramUsers{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		createData := entities.TelegramUsers{
			TelegramID: userData.Message.From.ID,
			FirstName:  userData.Message.From.FirstName,
			LastName:   userData.Message.From.LastName,
			Username:   userData.Message.From.UserName,
			Type:       userData.Message.Chat.Type,
		}
		queryDBByTelegramID := QueryExistTelegramID(userData.Message.From.ID)
		r.db.Raw(queryDBByTelegramID).Find(&countExistTelegramUser)
		if countExistTelegramUser.Count > 0 {
			msg = fmt.Sprintf("Tên Telegram này đã tồn tại hệ thống, TelegramID: %d | Tên tài khoản: %s", userData.Message.From.ID, userData.Message.From.UserName)
			ch <- false
			return
		}
		err = r.db.Debug().Model(&entities.TelegramUsers{}).Create(&createData).Error
		if err != nil {
			msg = fmt.Sprintf("%s", err)
			ch <- false
			return
		}
		queryDBForCountAllUsers := QueryAllTelegramUsers()
		r.db.Raw(queryDBForCountAllUsers).Find(&countCurrentUsers)
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		// Gửi tin Telegram thông báo
		telegramMsg := url.QueryEscape(ParamTelegramSendTextWelcome(userData.Message.From.UserName, int64(countCurrentUsers.Count)))
		replyMarkup := url.QueryEscape(ParamTelegramSendReplyMarkupWelcome())
		telegramParams.ChatID = userData.Message.From.ID
		telegramParams.Text = telegramMsg
		telegramParams.ReplyMarkup = replyMarkup
		getTelegramMsgUrl := config.GetTelegramMeEnglishSendMsgUrlConfig(telegramParams)
		return true, getTelegramMsgUrl
	}
	if msg == "" {
		msg = fmt.Sprintf("Lỗi không tạo đươc tài khoản")
	}
	telegramParams.Text = msg
	return false, config.GetTelegramMeEnglishSendMsgUrlConfig(telegramParams)
}
