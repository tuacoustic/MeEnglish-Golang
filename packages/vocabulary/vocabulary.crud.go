package vocabulary

import (
	"github.com/jinzhu/gorm"
)

type repositoryVocabularyCRUD struct {
	db *gorm.DB
}

func NewRepositoryVocabularyCRUD(db *gorm.DB) *repositoryVocabularyCRUD {
	return &repositoryVocabularyCRUD{db}
}
