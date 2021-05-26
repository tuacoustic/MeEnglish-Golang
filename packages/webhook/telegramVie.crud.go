package webhook

import (
	"github.com/jinzhu/gorm"
)

type repositoryTelegramVieCRUD struct {
	db *gorm.DB
}

func NewRepositoryTelegramVieCRUD(db *gorm.DB) *repositoryTelegramVieCRUD {
	return &repositoryTelegramVieCRUD{db}
}

func (r *repositoryTelegramVieCRUD) StudyNowVie(userData TelegramRespJSON) (bool, string) {

	return true, ""
}
