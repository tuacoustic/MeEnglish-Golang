package entities

import "time"

type AnswerKey struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	TelegramID   uint64    `json:"telegram_id"`
	VocabularyID uint64    `json:"vocabulary_id"`
	ExpiredAt    int64     `json:"expired_at"`
	Answer       string    `json:"answer"`
	AnswerType   string    `gorm:"size:20;default:'button'" json:"answer_type"`
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

type CountAnswerKey struct {
	Count uint64 `json:"count"`
}
