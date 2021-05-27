package webhook

import (
	"fmt"
	"me-english/entities"
	"me-english/utils/channels"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jinzhu/gorm"
)

type repositoryTelegramVieCRUD struct {
	db *gorm.DB
}

func NewRepositoryTelegramVieCRUD(db *gorm.DB) *repositoryTelegramVieCRUD {
	return &repositoryTelegramVieCRUD{db}
}

// Gửi về Group học đầu tiên
func (r *repositoryTelegramVieCRUD) StudyNowVie(userData TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) {
	var err error

	studyCommand := entities.GetTelegramStudyCommand{}
	countStudyCommand := entities.CountTelegramStudyCommand{}
	listsVocab := []entities.GetStudyVocab{}
	var textArr []string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// Count
		queryCountStudyCommand := QueryTelegramStudyCommand()
		r.db.Raw(queryCountStudyCommand).Find(&countStudyCommand)
		// console.Info(studyCommand)
		if countStudyCommand.Count > 0 {
			// Tìm kiếm khoá học gần nhất -> send về page khoá học đó
			err = r.db.Debug().Model(&entities.GetTelegramStudyCommand{}).Order("id desc").First(&studyCommand).Error
			if err != nil {
				msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
				ch <- false
				return
			}
			ch <- true
			return
		}
		// Ban đầu chưa có từ vựng học -> lấy tử vựng từ group 1
		err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", "1").Find(&listsVocab).Error
		if err != nil {
			msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
			ch <- false
			return
		}
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		if countStudyCommand.Count > 0 {
			return true, "", replyMarkup
		}
		for index, vocab := range listsVocab {
			if index%2 != 0 {
				secondVocabForEach := fmt.Sprintf("﹒%s {/%s}\n", vocab.Word, vocab.Word)
				textArr = append(textArr, secondVocabForEach)
			} else {
				firstVocabForEach := fmt.Sprintf("%s {/%s}", vocab.Word, vocab.Word)
				textArr = append(textArr, firstVocabForEach)
			}
		}
		textArrToString := strings.Join(textArr, "")
		text := StudyNowVie(1, textArrToString)
		replyMarkup = StudyNowVie_Reply
		return true, text, replyMarkup
	}
	return true, msg, replyMarkup
}
