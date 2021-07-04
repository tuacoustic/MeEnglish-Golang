package entities

import "time"

type StudyVocabLists struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	TelegramID   uint64    `json:"telegram_id"`
	VocabularyID uint64    `json:"vocabulary_id"`
	Active       bool      `gorm:"default:false" json:"active"`
	AwlGroupID   uint64    `json:"academic_group_id"`                   // ID của từ vựng
	PageNumber   uint64    `gorm:"size:2;default:1" json:"page_number"` // Limit 15 theo page
	Score        uint64    `gorm:"size:2;default:0" json:"score"`       // Số điểm
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
