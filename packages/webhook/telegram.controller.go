package webhook

import (
	"log"
	"me-english/database"
	"me-english/telegram"
	"me-english/utils/config"
	"me-english/utils/console"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ConnectWebhook() {
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_TOKEN_MEENGLISH)
	if err != nil {
		return
	}
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	var telegramPushWB TelegramRespJSON
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		telegramPushWB.Message = TelegramRespMessageJSON{
			MessageID: uint64(update.Message.Chat.ID),
			Text:      update.Message.Text,
			Date:      update.Message.Date,
		}
		telegramPushWB.Message.Chat.Type = update.Message.Chat.Type
		telegramPushWB.Message.From = TelegramRespMessageFromJSON{
			ID:        uint64(update.Message.From.ID),
			IsBot:     update.Message.From.IsBot,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
			UserName:  update.Message.From.UserName,
		}
		TelegramPushWebhook(telegramPushWB)
	}
	return
}

func TelegramPushWebhook(telegramPushWB TelegramRespJSON) {
	var commandFlag bool = false
	isCommand := string(telegramPushWB.Message.Text[0])
	if isCommand == Command_Handling.StartBot {
		commandFlag = true
	}
	db, err := database.MysqlConnect()
	if err != nil {
		return
	}
	defer db.Close()
	bot, err := telegram.ConnectBot()
	if err != nil {
		return
	}
	repo := NewRepositoryTelegramCRUD(db)
	GetStudyNowVie := NewRepositoryTelegramVieCRUD(db)
	switch commandFlag {
	case true: // Dùng lệnh
		// Xử lý khi khách nhập start -> Tạo User vào Database
		if telegramPushWB.Message.Text == "/start" {
			func(telegramRepo TelegramRepository) {
				status, text, replyMarkup := telegramRepo.CreateUser(telegramPushWB)
				if status == true {
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					msg.ParseMode = telegramParams.ParseMode
					msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
				msg.ParseMode = telegramParams.ParseMode
				bot.Send(msg)
				return
			}(repo)
		} else {
			// Tìm kiếm từ vựng
			func(telegramVieRepo TelegramVieRepository) {
				status, text, replyMarkup := telegramVieRepo.FindVocab(telegramPushWB)
				if status == true {
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					msg.ParseMode = telegramParams.ParseMode
					msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
				msg.ParseMode = telegramParams.ParseMode
				msg.ReplyMarkup = replyMarkup
				bot.Send(msg)
				return
			}(GetStudyNowVie)
		}
		break
	case false: // Không dùng lệnh
		switch strings.ToLower(telegramPushWB.Message.Text) {
		case Command_Handling.GetStudyNowVie:
			func(telegramVieRepo TelegramVieRepository) {
				status, text, replyMarkup := telegramVieRepo.GetStudyNowVie(telegramPushWB)
				if status == true {
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					msg.ParseMode = telegramParams.ParseMode
					msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
				msg.ParseMode = telegramParams.ParseMode
				msg.ReplyMarkup = replyMarkup
				bot.Send(msg)
				return
			}(GetStudyNowVie)
			break
		case Command_Handling.AutoRemindVie:
			console.Info("Nhắc học tự động")
			break
		case Command_Handling.InstructionVie:
			console.Info("Hướng dẫn")
			break
		case Command_Handling.SupportVie:
			msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), Support_Text)
			msg.ParseMode = telegramParams.ParseMode
			bot.Send(msg)
			break
		case Command_Handling.DevelopVie:
			console.Info("Cùng phát triển")
			break
		case Command_Handling.DonateVie:
			msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), Donate_Text)
			msg.ParseMode = telegramParams.ParseMode
			bot.Send(msg)
			break
		case Command_Handling.BackHome:
			func(telegramVieRepo TelegramVieRepository) {
				status, text, replyMarkup := telegramVieRepo.BackHomePage(telegramPushWB)
				if status == true {
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					msg.ParseMode = telegramParams.ParseMode
					msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
				msg.ParseMode = telegramParams.ParseMode
				msg.ReplyMarkup = replyMarkup
				bot.Send(msg)
				return
			}(GetStudyNowVie)
			break
		case Command_Handling.Continue:
			func(telegramVieRepo TelegramVieRepository) {
				status, text, replyMarkup := telegramVieRepo.GetStudyNowVie(telegramPushWB)
				if status == true {
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					msg.ParseMode = telegramParams.ParseMode
					msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
				msg.ParseMode = telegramParams.ParseMode
				msg.ReplyMarkup = replyMarkup
				bot.Send(msg)
				return
			}(GetStudyNowVie)
			break
		default:
			initCondition := strings.ToLower(telegramPushWB.Message.Text)
			switch {
			case strings.Contains(initCondition[0:3], Command_Handling.OnCurrentPage) == true || strings.Contains(initCondition, ">gr") == true:
				break
			case strings.Contains(initCondition[0:2], Command_Handling.QueryPage) == true:
				func(telegramVieRepo TelegramVieRepository) {
					status, text, replyMarkup := telegramVieRepo.GetVocabByGroupPageVie(telegramPushWB)
					if status == true {
						msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
						msg.ParseMode = telegramParams.ParseMode
						msg.ReplyMarkup = replyMarkup
						bot.Send(msg)
						return
					}
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					// msg.ParseMode = telegramParams.ParseMode
					// msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}(GetStudyNowVie)
				break
			case strings.Contains(initCondition[0:2], Command_Handling.QueryGroup) == true:
				func(telegramVieRepo TelegramVieRepository) {
					status, text, replyMarkup := telegramVieRepo.GetVocabByGroupVie(telegramPushWB)
					if status == true {
						msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
						msg.ParseMode = telegramParams.ParseMode
						msg.ReplyMarkup = replyMarkup
						bot.Send(msg)
						return
					}
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					// msg.ParseMode = telegramParams.ParseMode
					// msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}(GetStudyNowVie)
				break
			case strings.Contains(initCondition[0:16], Command_Handling.GroupStudy) == true:
				func(telegramVieRepo TelegramVieRepository) {
					status, text, replyMarkup := telegramVieRepo.GroupStudy(telegramPushWB)
					if status == true {
						msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
						msg.ParseMode = telegramParams.ParseMode
						msg.ReplyMarkup = replyMarkup
						bot.Send(msg)
						return
					}
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					// msg.ParseMode = telegramParams.ParseMode
					// msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}(GetStudyNowVie)
				break
			}
		}
	}
	return
}
