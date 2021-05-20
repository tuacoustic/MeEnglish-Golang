package entities

import "time"

type TelegramUsers struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CustomerID uint64    `gorm:"unique" json:"customer_id"`
	TelegramID uint64    `gorm:"unique" json:"telegram_id"`
	FirstName  string    `gorm:"size:100" json:"first_name"`
	LastName   string    `gorm:"size:100" json:"last_name"`
	Type       string    `gorm:"50" json:"type"`
	ActiveCode bool      `gorm:"default:false" json:"active_code"`
	CreatedAt  time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

type CountTelegramUsers struct {
	Count uint32
}
