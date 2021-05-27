package entities

func (GetTelegramStudyCommand) TableName() string {
	return "telegram_study_command"
}

func (GetStudyVocab) TableName() string {
	return "vocabulary"
}
