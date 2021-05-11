package entities

import "time"

// Vocabulary
type Vocabulary struct {
	ID               uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Word             string    `gorm:"size:255;not null;unique" json:"word"`
	Provider         string    `gorm:"size:100;not null;default:'Oxford University Press'" json:"provider"`
	Language         string    `gorm:"size:20;not null;default:'es-us'" json:"language"`
	Tag              string    `gorm:"size:50;not null;default:'vocabulary'" json:"tag"`
	Vi               string    `gorm:"size:255" json:"vi"`
	Image            string    `gorm:"size:255;default:'https://st4.depositphotos.com/14953852/24787/v/600/depositphotos_247872612-stock-illustration-no-image-available-icon-vector.jpg'" json:"image"`
	Definition       string    `gorm:"size:255" json:"definition"`
	Example          string    `gorm:"type:text" json:"Example"`
	AudioFile        string    `gorm:"size:255" json:"audio_file"`
	Dialects         string    `gorm:"size:100:default:'American English'" json:"dialects"`
	PhoneticSpelling string    `gorm:"size:100" json:"phonetic_spelling"`
	CreatedAt        time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
