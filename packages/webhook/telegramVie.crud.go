package webhook

import (
	"fmt"
	"me-english/entities"
	"me-english/utils/channels"
	"strconv"
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
func (r *repositoryTelegramVieCRUD) GetStudyNowVie(userData TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) {
	var err error

	studyCommand := entities.GetTelegramStudyCommand{}
	countStudyCommand := entities.CountTelegramStudyCommand{}
	countVocabByGroup := entities.CountVocabulary{}
	listsVocab := []entities.GetStudyVocab{}
	var textArr []string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// Count
		queryCountStudyCommand := QueryTelegramStudyGroupCommand(userData.Message.From.ID, EnumStudyCommand.GetCommand)
		r.db.Raw(queryCountStudyCommand).Find(&countStudyCommand)
		if countStudyCommand.Count > 0 {
			if countStudyCommand.Count > 1 { // Logic hơn 1 Group active
				getUpdateLatest := entities.GetTelegramStudyCommand{}
				r.db.Model(&entities.GetTelegramStudyCommand{}).Select("id, awl_group_id").Where("telegram_id = ? and active = true", userData.Message.From.ID).Order("created_at desc").Find(&getUpdateLatest)
				if err != nil {
					msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
					ch <- false
					return
				}
				// Update != latest active = false
				err = r.db.Debug().Model(&entities.GetTelegramStudyCommand{}).Where("id != ? and telegram_id = ?", getUpdateLatest.ID, userData.Message.From.ID).Update("active", "false").Error
				if err != nil {
					msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
					ch <- false
					return
				}
				queryCountVocab := QueryCountVocabByGroup(getUpdateLatest.AwlGroupID)
				r.db.Raw(queryCountVocab).Find(&countVocabByGroup)
				if int(countVocabByGroup.Count) <= Limit_GetVocab {
					err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", getUpdateLatest.AwlGroupID).Limit(Limit_GetVocab).Find(&listsVocab).Error
					if err != nil {
						msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
						ch <- false
						return
					}
					studyCommand = getUpdateLatest
					ch <- true
					return
				}
				err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", getUpdateLatest.AwlGroupID).Limit(Limit_GetVocab).Offset(0).Find(&listsVocab).Error
				if err != nil {
					msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
					ch <- false
					return
				}

				// Đưa cho thằng studyCommand reuse
				ch <- true
				return
			}
			// Tìm kiếm khoá học gần nhất -> send về page khoá học đó
			err = r.db.Debug().Model(&entities.GetTelegramStudyCommand{}).Select("awl_group_id").Where("telegram_id = ? and active = true", userData.Message.From.ID).Order("created_at desc").First(&studyCommand).Error
			if err != nil {
				msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
				ch <- false
				return
			}
			queryCountVocab := QueryCountVocabByGroup(studyCommand.AwlGroupID)
			r.db.Raw(queryCountVocab).Find(&countVocabByGroup)
			if int(countVocabByGroup.Count) <= Limit_GetVocab {
				err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", studyCommand.AwlGroupID).Limit(Limit_GetVocab).Find(&listsVocab).Error
				if err != nil {
					msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
					ch <- false
					return
				}
				ch <- true
				return
			}
			err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", studyCommand.AwlGroupID).Limit(Limit_GetVocab).Offset(0).Find(&listsVocab).Error
			if err != nil {
				msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
				ch <- false
				return
			}
			ch <- true
			return
		}
		// Ban đầu chưa có từ vựng học -> lấy tử vựng từ group 1
		// Đến số lượng từ Group 1 -> Chia paginate
		queryCountVocab := QueryCountVocabByGroup(1)
		r.db.Raw(queryCountVocab).Find(&countVocabByGroup)
		err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", "1").Limit(Limit_GetVocab).Offset(0).Find(&listsVocab).Error
		if err != nil {
			msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
			ch <- false
			return
		}
		// Create Study Command
		createStudyCommand := entities.TelegramStudyCommand{
			TelegramID: userData.Message.From.ID,
			Username:   userData.Message.From.UserName,
			Command:    Command_GetGroup,
			TextInput:  userData.Message.Text,
			AwlGroupID: 1,
			Active:     true,
			Timestamp:  userData.Message.Date,
		}
		r.db.Model(&entities.TelegramStudyCommand{}).Create(&createStudyCommand)
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		replyMarkup = StudyNowVieReplyDefault(studyCommand.AwlGroupID)
		if countStudyCommand.Count > 0 {
			if len(listsVocab) == 0 {
				text := GetStudyNowNullVocabVie(studyCommand.AwlGroupID)
				return false, text, replyMarkup
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
			text := GetStudyNowVie(studyCommand.AwlGroupID, 1, textArrToString)
			replyMarkup = StudyNowVieReply(studyCommand.AwlGroupID, 1, countVocabByGroup.Count)
			return true, text, replyMarkup
		}
		if len(listsVocab) == 0 {
			text := GetStudyNowNullVocabVie(studyCommand.AwlGroupID)
			return false, text, replyMarkup
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
		text := GetStudyNowVie(1, 1, textArrToString)
		replyMarkup = StudyNowVieReply(1, 1, countVocabByGroup.Count)
		return true, text, replyMarkup
	}
	return true, msg, replyMarkup
}

func (r *repositoryTelegramVieCRUD) GetVocabByGroupPageVie(userData TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) {
	var err error

	studyCommand := entities.GetTelegramStudyCommand{}
	countStudyCommand := entities.CountTelegramStudyCommand{}
	countVocabByGroup := entities.CountVocabulary{}
	listsVocab := []entities.GetStudyVocab{}
	numberOfPage := userData.Message.Text[2:3]
	sendNumberOfPage, err := strconv.ParseUint(numberOfPage, 10, 32)
	var textArr []string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup
	if err != nil {
		return false, "Bạn gửi sai cú pháp! Vui lòng thử lại nhé 😉", replyMarkup
	}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		queryCountStudyCommand := QueryTelegramStudyGroupCommand(userData.Message.From.ID, EnumStudyCommand.GetCommand)
		r.db.Raw(queryCountStudyCommand).Find(&countStudyCommand)
		// Logic hơn 1 Group Active
		if countStudyCommand.Count > 1 {
			getUpdateLatest := entities.GetTelegramStudyCommand{}
			r.db.Model(&entities.GetTelegramStudyCommand{}).Select("id, awl_group_id").Where("telegram_id = ? and active = true", userData.Message.From.ID).Order("created_at desc").Find(&getUpdateLatest)
			if err != nil {
				msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
				ch <- false
				return
			}
			// Update != latest active = false
			err = r.db.Debug().Model(&entities.GetTelegramStudyCommand{}).Where("id != ? and telegram_id = ?", getUpdateLatest.ID, userData.Message.From.ID).Update("active", "false").Error
			if err != nil {
				msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
				ch <- false
				return
			}
			queryCountVocab := QueryCountVocabByGroup(getUpdateLatest.AwlGroupID)
			r.db.Raw(queryCountVocab).Find(&countVocabByGroup)
			if int(countVocabByGroup.Count) <= Limit_GetVocab {
				err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", getUpdateLatest.AwlGroupID).Limit(Limit_GetVocab).Find(&listsVocab).Error
				if err != nil {
					msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
					ch <- false
					return
				}
				studyCommand = getUpdateLatest
				ch <- true
				return
			}
			err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", getUpdateLatest.AwlGroupID).Limit(Limit_GetVocab).Offset(0).Find(&listsVocab).Error
			if err != nil {
				msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
				ch <- false
				return
			}
			ch <- true
			return
		}
		// Tìm kiếm khoá học gần nhất -> send về page khoá học đó
		err = r.db.Debug().Model(&entities.GetTelegramStudyCommand{}).Select("awl_group_id").Where("telegram_id = ? and active = true", userData.Message.From.ID).Order("created_at desc").First(&studyCommand).Error
		if err != nil {
			msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
			ch <- false
			return
		}
		queryCountVocab := QueryCountVocabByGroup(studyCommand.AwlGroupID)
		r.db.Raw(queryCountVocab).Find(&countVocabByGroup)
		if int(countVocabByGroup.Count) <= Limit_GetVocab {
			err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", studyCommand.AwlGroupID).Limit(Limit_GetVocab).Find(&listsVocab).Error
			if err != nil {
				msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
				ch <- false
				return
			}
			ch <- true
			return
		}
		offset := (int(sendNumberOfPage) - 1) * Limit_GetVocab
		err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", studyCommand.AwlGroupID).Limit(Limit_GetVocab).Offset(offset).Find(&listsVocab).Error
		if err != nil {
			msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
			ch <- false
			return
		}
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		replyMarkup = StudyNowVieReplyDefault(studyCommand.AwlGroupID)
		// Check numberOfPage
		var maxPaginationNumber int = int(countVocabByGroup.Count / uint32(Limit_GetVocab))
		if float64(countVocabByGroup.Count)/float64(Limit_GetVocab) > float64(maxPaginationNumber) {
			maxPaginationNumber = maxPaginationNumber + 1
		}
		if int(sendNumberOfPage) > maxPaginationNumber {
			return false, "Bạn gửi sai cú pháp! Vui lòng thử lại nhé 😉", replyMarkup
		}
		if countStudyCommand.Count > 0 {
			if len(listsVocab) == 0 {
				text := GetStudyNowNullVocabVie(studyCommand.AwlGroupID)
				return false, text, replyMarkup
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
			text := GetStudyNowVie(studyCommand.AwlGroupID, sendNumberOfPage, textArrToString)
			replyMarkup = StudyNowVieReply(studyCommand.AwlGroupID, uint64(sendNumberOfPage), countVocabByGroup.Count)
			return true, text, replyMarkup
		}
		if len(listsVocab) == 0 {
			text := GetStudyNowNullVocabVie(studyCommand.AwlGroupID)
			return false, text, replyMarkup
		}
	}
	return false, msg, replyMarkup
}

func (r *repositoryTelegramVieCRUD) GetVocabByGroupVie(userData TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) {
	var err error

	// studyCommand := entities.GetTelegramStudyCommand{}
	var replyMarkup tgbotapi.ReplyKeyboardMarkup
	// countExistStudyCommand := entities.CountTelegramStudyCommand{}
	var numberOfGroup string
	countVocabByGroup := entities.CountVocabulary{}
	listsVocab := []entities.GetStudyVocab{}
	var textArr []string
	if len(userData.Message.Text) >= 4 {
		numberOfGroup = userData.Message.Text[2:4]
	} else {
		numberOfGroup = userData.Message.Text[2:3]
	}
	sendNumberOfGroup, err := strconv.ParseUint(numberOfGroup, 10, 32)
	if err != nil {
		return false, "Bạn gửi sai cú pháp! Vui lòng thử lại nhé 😉", replyMarkup
	}
	// studyCommand := entities.GetTelegramStudyCommand{}
	groupInVocabExist := entities.CountVocabulary{}
	// countStudyCommand := entities.CountTelegramStudyCommand{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// Tìm group trong list vocab
		r.db.Debug().Table("vocabulary").Select("awl_group_id as count").Where("awl_group_id = ?", sendNumberOfGroup).Group("awl_group_id").First(&groupInVocabExist)

		// Tạo GET_GROUP trong study command
		// Check exist
		if groupInVocabExist.Count >= 1 {
			// Count Vocab
			queryCountVocab := QueryCountVocabByGroup(sendNumberOfGroup)
			r.db.Debug().Raw(queryCountVocab).Find(&countVocabByGroup)
			if int(countVocabByGroup.Count) <= Limit_GetVocab {
				err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", sendNumberOfGroup).Limit(Limit_GetVocab).Find(&listsVocab).Error
				if err != nil {
					msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
					ch <- false
					return
				}
				// studyCommand = getUpdateLatest
				ch <- true
				return
			}
			err = r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", sendNumberOfGroup).Limit(Limit_GetVocab).Offset(0).Find(&listsVocab).Error
			if err != nil {
				msg = "Oops Lỗi rồi, bạn thử lại sau nhé 😉"
				ch <- false
				return
			}
			ch <- true
			return
		}
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		replyMarkup = StudyNowVieReplyDefault(sendNumberOfGroup)
		if groupInVocabExist.Count >= 1 {
			if len(listsVocab) == 0 {
				text := GetStudyNowNullVocabVie(sendNumberOfGroup)
				return false, text, replyMarkup
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
			text := GetStudyNowVie(sendNumberOfGroup, 1, textArrToString)
			replyMarkup = StudyNowVieReply(sendNumberOfGroup, 1, countVocabByGroup.Count)
			return true, text, replyMarkup
		}
		text := GetStudyNowNullVocabVie(sendNumberOfGroup)
		return false, text, replyMarkup
	}
	return false, msg, replyMarkup
}

func (r *repositoryTelegramVieCRUD) BackHomePage(userData TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) {
	var replyMarkup tgbotapi.ReplyKeyboardMarkup
	countStudyCommand := entities.CountTelegramStudyCommand{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		queryCountStudyCommand := QueryTelegramStudyGroupCommand(userData.Message.From.ID, EnumStudyCommand.StudyCommand)
		r.db.Raw(queryCountStudyCommand).Find(&countStudyCommand)
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		if countStudyCommand.Count >= 1 {
			replyMarkup = Back_Home_Reply
			text := "Bạn đã trở về trang chủ"
			return true, text, replyMarkup
		}
		replyMarkup = Home_Reply
		text := "Bạn ở trang chủ"
		return true, text, replyMarkup
	}
	return false, msg, replyMarkup
}

func (r *repositoryTelegramVieCRUD) GroupStudy(userData TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) {
	var err error
	var replyMarkup tgbotapi.ReplyKeyboardMarkup

	var numberOfGroupStudy string
	if len(userData.Message.Text) >= 19 {
		numberOfGroupStudy = userData.Message.Text[17:19]
	} else {
		numberOfGroupStudy = userData.Message.Text[17:18]
	}

	sendNumberOfGroup, err := strconv.ParseInt(numberOfGroupStudy, 10, 32)
	if err != nil {
		return false, "Bạn gửi sai cú pháp! Vui lòng thử lại nhé 😉", replyMarkup
	}
	if sendNumberOfGroup > 10 {
		return false, "Bạn gửi sai cú pháp! Vui lòng thử lại nhé 😉", replyMarkup
	}

	done := make(chan bool)
	go func(ch chan bool) {
		defer close(ch)
		// Find Study Group
	}(done)
	return false, msg, replyMarkup
}

func (r *repositoryTelegramVieCRUD) FindVocab(userData TelegramRespJSON) (bool, string, tgbotapi.ReplyKeyboardMarkup) {
	var replyMarkup tgbotapi.ReplyKeyboardMarkup
	vocab := userData.Message.Text[1:]
	getVocabDetail := entities.FindVocab{}
	listsVocab := []entities.GetStudyVocab{}
	countVocabByGroup := entities.CountVocabulary{}
	// var text string
	var textArr []string
	// var sendNumberOfGroup uint64
	done := make(chan bool)
	go func(ch chan bool) {
		defer close(ch)
		r.db.Debug().Table("vocabulary").Where("word like ?", vocab).First(&getVocabDetail)
		if getVocabDetail.Word == "" {
			msg = "Không tìm thấy từ vựng!"
			ch <- false
			return
		}
		offset := (int(getVocabDetail.PageNumber) - 1) * Limit_GetVocab
		r.db.Debug().Model(&entities.GetStudyVocab{}).Where("awl_group_id = ?", getVocabDetail.AwlGroupID).Limit(Limit_GetVocab).Offset(offset).Find(&listsVocab)
		queryCountVocab := QueryCountVocabByGroup(getVocabDetail.AwlGroupID)
		r.db.Debug().Raw(queryCountVocab).Find(&countVocabByGroup)
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		if getVocabDetail.Word != "" {
			if len(listsVocab) == 0 {
				text := "Không tìm thấy từ vựng!"
				return false, text, replyMarkup
			}
			for index, vocab := range listsVocab {
				if index%2 != 0 {
					if vocab.Word == getVocabDetail.Word {
						secondVocabForEach := fmt.Sprintf("﹒🔍 %s {/%s}\n", vocab.Word, vocab.Word)
						textArr = append(textArr, secondVocabForEach)
						continue
					}
					secondVocabForEach := fmt.Sprintf("﹒%s {/%s}\n", vocab.Word, vocab.Word)
					textArr = append(textArr, secondVocabForEach)
				} else {
					if vocab.Word == getVocabDetail.Word {
						firstVocabForEach := fmt.Sprintf("🔍 %s {/%s}", vocab.Word, vocab.Word)
						textArr = append(textArr, firstVocabForEach)
						continue
					}
					firstVocabForEach := fmt.Sprintf("%s {/%s}", vocab.Word, vocab.Word)
					textArr = append(textArr, firstVocabForEach)
				}
			}
			text := VocabDetailText(1, getVocabDetail, textArr)
			sendNumberOfGroup := getVocabDetail.AwlGroupID
			replyMarkup = StudyNowVieReply(sendNumberOfGroup, 1, countVocabByGroup.Count)
			return true, text, replyMarkup
		}
	}
	return false, msg, replyMarkup
}
