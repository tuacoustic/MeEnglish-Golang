package product

import (
	"github.com/jinzhu/gorm"
)

type repositoryProductCRUD struct {
	db *gorm.DB
}

func NewRepositoryProductCRUD(db *gorm.DB) *repositoryProductCRUD {
	return &repositoryProductCRUD{db}
}
