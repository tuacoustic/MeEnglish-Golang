package webhook

import (
	"fmt"
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
	if isCommand == Command_Handling.BotCommand {
		commandFlag = true
	}
	db, err := database.MysqlConnect()
	if err != nil {
		return
	}
	// defer db.Close()
	bot, err := telegram.ConnectBot()
	if err != nil {
		return
	}
	repo := NewRepositoryTelegramCRUD(db)
	GetStudyNowVie := NewRepositoryTelegramVieCRUD(db)
	switch commandFlag {
	// Xử lý khi khách nhập start -> Tạo User vào Database
	case true: // Dùng lệnh
		initCondition := strings.ToLower(telegramPushWB.Message.Text)
		switch {
		case strings.Contains(initCondition[0:6], Command_Handling.StartBot) == true:
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
			break
		case strings.Contains(initCondition[0:7], Command_Handling.GetAudio) == true:
			func(telegramVieRepo TelegramVieRepository) {
				status, filePath, errMsg := telegramVieRepo.FindAudio(telegramPushWB)
				if status == true {
					var msg tgbotapi.AudioConfig
					msg = tgbotapi.NewAudioShare(int64(telegramPushWB.Message.From.ID), filePath)
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), fmt.Sprintf("%s", errMsg))
				bot.Send(msg)
				return
			}(GetStudyNowVie)
			break
		case strings.Contains(initCondition[0:7], Command_Handling.GetImage) == true:
			func(telegramVieRepo TelegramVieRepository) {
				status, filePath, errMsg := telegramVieRepo.FindImage(telegramPushWB)
				if status == true {
					console.Info(filePath)
					var msg tgbotapi.PhotoConfig
					msg = tgbotapi.NewPhotoShare(int64(telegramPushWB.Message.From.ID), filePath)
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), fmt.Sprintf("%s", errMsg))
				bot.Send(msg)
				return
			}(GetStudyNowVie)
			break
		default:
			// Tìm kiếm từ vựng
			func(telegramVieRepo TelegramVieRepository) {
				status, text, _ := telegramVieRepo.FindVocab(telegramPushWB)
				if status == true {
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					msg.ParseMode = telegramParams.ParseMode
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
				msg.ParseMode = telegramParams.ParseMode
				bot.Send(msg)
				return
			}(GetStudyNowVie)
			break
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
		case Command_Handling.Continue:
			func(telegramVieRepo TelegramVieRepository) {
				status, text, _, replyMarkup := telegramVieRepo.HandleTrueAnswer(telegramPushWB)
				if status == true {
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					msg.ParseMode = telegramParams.ParseMode
					msg.ReplyMarkup = replyMarkup
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
				bot.Send(msg)
				return
			}(GetStudyNowVie)
			break
		case Command_Handling.SelectGroup:
			fallthrough
		case Command_Handling.AnotherGroup:
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
		case Command_Handling.Suggestion:
			func(telegramVieRepo TelegramVieRepository) {
				status, text, _ := telegramVieRepo.ShowAnswer(telegramPushWB)
				if status == true {
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					msg.ParseMode = telegramParams.ParseMode
					bot.Send(msg)
					func(telegramVieRepo TelegramVieRepository) {
						status, text, _, replyMarkup := telegramVieRepo.HandleTrueAnswer(telegramPushWB)
						if status == true {
							msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
							msg.ParseMode = telegramParams.ParseMode
							msg.ReplyMarkup = replyMarkup
							bot.Send(msg)
							return
						}
						msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
						bot.Send(msg)
					}(GetStudyNowVie)
					return
				}
				msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
				bot.Send(msg)
			}(GetStudyNowVie)
			break
		default:
			initCondition := strings.ToLower(telegramPushWB.Message.Text)
			console.Info(initCondition)
			var lengthText uint32
			var sliceTextByLength string
			lengthText = uint32(len(initCondition))
			if lengthText == 16 {
				sliceTextByLength = initCondition[0:16]
			} else if lengthText == 6 {
				sliceTextByLength = initCondition[0:6]
			} else if lengthText == 3 {
				sliceTextByLength = initCondition[0:3]
			} else if lengthText == 2 {
				sliceTextByLength = initCondition[0:2]
			} else {
				sliceTextByLength = initCondition
			}
			switch {
			case strings.Contains(sliceTextByLength, Command_Handling.OnCurrentPage) == true || strings.Contains(initCondition, ">gr") == true:
				break
			case strings.Contains(sliceTextByLength, Command_Handling.QueryPage) == true && (lengthText <= 4 && lengthText >= 3):
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
					bot.Send(msg)
					return
				}(GetStudyNowVie)
				break
			case strings.Contains(sliceTextByLength, Command_Handling.QueryGroup) == true && (lengthText <= 4 && lengthText >= 3):
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
					bot.Send(msg)
					return
				}(GetStudyNowVie)
				break
			case strings.Contains(sliceTextByLength, Command_Handling.AnswerButton) == true && lengthText == 8:
				func(telegramVieRepo TelegramVieRepository) {
					status, text, replyMarkup := telegramVieRepo.AnswerQuestionButton(telegramPushWB)
					if status == true {
						msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
						msg.ParseMode = telegramParams.ParseMode
						msg.ReplyMarkup = replyMarkup
						bot.Send(msg)
						// Send Next Question
						func(telegramVieRepo TelegramVieRepository) {
							status, text, _, replyMarkup := telegramVieRepo.HandleTrueAnswer(telegramPushWB)
							if status == true {
								msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
								msg.ParseMode = telegramParams.ParseMode
								msg.ReplyMarkup = replyMarkup
								bot.Send(msg)
								return
							}
							msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
							bot.Send(msg)
							return
						}(GetStudyNowVie)
						return
					}
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					bot.Send(msg)
					return
				}(GetStudyNowVie)
				break
			case strings.Contains(sliceTextByLength, Command_Handling.GroupStudy) == true && (lengthText <= 19 && lengthText >= 18):
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
					bot.Send(msg)
					return
				}(GetStudyNowVie)
				break
			case strings.Contains(sliceTextByLength, Command_Handling.TrueAnswer) == true && lengthText == 16:
				func(telegramVieRepo TelegramVieRepository) {
					status, text, _, replyMarkup := telegramVieRepo.HandleTrueAnswer(telegramPushWB)
					if status == true {
						msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
						msg.ParseMode = telegramParams.ParseMode
						msg.ReplyMarkup = replyMarkup
						bot.Send(msg)
						return
					}
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					bot.Send(msg)
				}(GetStudyNowVie)
				break
			default:
				func(telegramVieRepo TelegramVieRepository) {
					status, text, replyMarkup := telegramVieRepo.AnswerQuestionByText(telegramPushWB)
					if status == true {
						msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
						msg.ParseMode = telegramParams.ParseMode
						msg.ReplyMarkup = replyMarkup
						bot.Send(msg)
						// Send Next Question
						func(telegramVieRepo TelegramVieRepository) {
							status, text, _, replyMarkup := telegramVieRepo.HandleTrueAnswer(telegramPushWB)
							if status == true {
								msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
								msg.ParseMode = telegramParams.ParseMode
								msg.ReplyMarkup = replyMarkup
								bot.Send(msg)
								return
							}
							msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
							bot.Send(msg)
							return
						}(GetStudyNowVie)
						return
					}
					msg := tgbotapi.NewMessage(int64(telegramPushWB.Message.From.ID), text)
					bot.Send(msg)
					return
				}(GetStudyNowVie)
				break
			}
		}
	}
	return
}
