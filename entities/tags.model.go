package entities

import "time"

// Thẻ
type Tag struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	NameTag   string    `gorm:"size:100;not null;unique" json:"name_tag"` // Tên thẻ
	DescTag   string    `gorm:"type:text" json:"desc_tag"`                // Mô tả thẻ
	ProductID uint64    `gorm:"not null" json:"product_id"`               // ID của sản phẩm
	StatusTag int32     `gorm:"default:0" json:"status_tag"`              // Trạng thái thẻ (0: False | 1: True)
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
