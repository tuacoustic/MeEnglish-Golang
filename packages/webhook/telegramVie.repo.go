package webhook

type TelegramVieRepository interface {
	// Vie
	StudyNowVie(TelegramRespJSON) (bool, string) // Status | url
}
