package webhook

import (
	"errors"
	"fmt"
	"me-english/entities"
	"me-english/utils/channels"

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
	var err error
	countEistTelegramUser := entities.CountTelegramUsers{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		createData := entities.TelegramUsers{
			TelegramID: userData.Message.From.ID,
			FirstName:  userData.Message.From.FirstName,
			LastName:   userData.Message.From.LastName,
			Type:       userData.Message.Chat.Type,
		}
		queryDBByTelegramID := QueryExistTelegramID(userData.Message.From.ID)
		r.db.Raw(queryDBByTelegramID).Find(&countEistTelegramUser)
		if countEistTelegramUser.Count > 0 {
			msg = "Tên Telegram này đã tồn tại hệ thống"
			ch <- false
			return
		}
		err = r.db.Debug().Model(&entities.TelegramUsers{}).Create(&createData).Error
		if err != nil {
			msg = fmt.Sprintf("%s", err)
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return true, nil
	}
	return false, errors.New(msg)
}
