package webhook

type TelegramRepository interface {
	CreateUser(TelegramRespJSON) (bool, string) // Status | url
}
