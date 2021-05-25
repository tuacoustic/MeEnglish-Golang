package webhook

import (
	"log"
	"me-english/database"
	"me-english/utils/config"
	sendReq "me-english/utils/sendRequest"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Connect() {
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_TOKEN_MEENGLISH)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	var telegramPushWB TelegramRespJSON
	var telegramMessageEntities []TelegramRespEntitiesJSON
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		telegramPushWB.Message = TelegramRespMessageJSON{
			MessageID: uint64(update.Message.Chat.ID),
			Text:      update.Message.Text,
		}
		telegramPushWB.Message.Chat.Type = update.Message.Chat.Type
		telegramPushWB.Message.From = TelegramRespMessageFromJSON{
			ID:        uint64(update.Message.From.ID),
			IsBot:     update.Message.From.IsBot,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
			UserName:  update.Message.From.UserName,
		}
		// Add detail outside loop
		for _, value := range *update.Message.Entities {
			telegramMessageEntities = append(telegramMessageEntities, TelegramRespEntitiesJSON{
				Type: value.Type,
			})
		}
		telegramPushWB.Message.Entities = telegramMessageEntities
		TelegramPushWebhook(telegramPushWB)
	}
}

func TelegramPushWebhook(telegramPushWB TelegramRespJSON) {
	var commandFlag bool = false
	for _, entityType := range telegramPushWB.Message.Entities {
		if entityType.Type == BOT_COMMAND {
			commandFlag = true
		}
	}
	db, err := database.MysqlConnect()
	if err != nil {
		return
	}
	repo := NewRepositoryTelegramCRUD(db)
	switch commandFlag {
	case true: // Dùng lệnh
		// Xử lý khi khách nhập start -> Tạo User vào Database
		func(telegramRepo TelegramRepository) {
			status, url := telegramRepo.CreateUser(telegramPushWB)
			if status == true {
				sendReq.PostRequestToTelegram(url, "GET", "")
				return
			}
			sendReq.PostRequestToTelegram(url, "GET", "")
			return
		}(repo)
		break
	case false: // Không dùng lệnh

		break
	}
	return
}
