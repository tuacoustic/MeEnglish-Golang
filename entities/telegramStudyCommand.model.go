package entities

import "time"

type TelegramStudyCommand struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CustomerID uint64    `json:"customer_id"`
	TelegramID uint64    `json:"telegram_id"`
	Username   string    `gorm:"size:100" json:"username"`
	Command    string    `gorm:"size:50" json:"command"`
	TextInput  string    `gorm:"size:50" json:"text_input"`
	AwlGroupID uint64    `json:"academic_group_id"`
	Active     bool      `gorm:"default:0" json:"active"`
	Timestamp  int       `gorm:"size:11" json:"timestamp"`
	CreatedAt  time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

type GetTelegramStudyCommand struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	CustomerID uint64 `json:"customer_id"`
	TelegramID uint64 `json:"telegram_id"`
	Username   string `gorm:"size:100" json:"username"`
	Command    string `gorm:"size:50" json:"command"`
	TextInput  string `gorm:"size:50" json:"text_input"`
	AwlGroupID uint64 `json:"awl_group_id"`
	Active     bool   `json:"active"`
}

type CountTelegramStudyCommand struct {
	Count uint32
}

// type GetIDTelegramStudyCommand struct {
// 	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
// }
