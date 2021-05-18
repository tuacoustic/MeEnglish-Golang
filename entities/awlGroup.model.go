package entities

import "time"

type AwlGroup struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:100;not null;default:'Group 1'" json:"title_group"`
	Desc      string    `gorm:"type:text" json:"desc_group"`
	Image     string    `gorm:"size:255;default:'https://st4.depositphotos.com/14953852/24787/v/600/depositphotos_247872612-stock-illustration-no-image-available-icon-vector.jpg'" json:"image"` // Hình ảnh giải thích chung
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

type GetAllAwlGroup struct {
	Title string `gorm:"size:100;not null;default:'Group 1'" json:"title_group"`
	Desc  string `gorm:"type:text;not null;default:'Đây là Group 1'" json:"desc_group"`
	Image string `gorm:"size:255;" json:"image"` // Hình ảnh giải thích chung
}
