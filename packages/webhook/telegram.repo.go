package webhook

type TelegramRepository interface {
	CreateUser(TelegramRespJSON) (bool, error)
}
