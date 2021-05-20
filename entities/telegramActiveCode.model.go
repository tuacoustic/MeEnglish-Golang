package entities

import "time"

type TelegramActiveCode struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CustomerID uint64    `json:"customerID"`
	TelegramID uint64    `json:"telegram_id"`
	ActiveCode string    `gorm:"size:6;" json:"active_code"`
	Active     bool      `gorm:"default:false" json:"active"`
	Expire     string    `gorm:"size:10;default:'1621496452'" json:"exp"`
	CreatedAt  time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
