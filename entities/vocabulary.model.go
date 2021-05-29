package entities

import "time"

// Vocabulary | noun | verb | adjective | adverb | phrasal
type Vocabulary struct {
	ID                  uint64    `gorm:"primary_key;auto_increment" json:"id"`
	AwlGroupID          uint64    `json:"academic_group_id"`                                                                                                                                                          // ID của từ vựng
	Word                string    `gorm:"size:255;not null;unique" json:"word"`                                                                                                                                       // Từ ngữ - Tiếng Anh
	Provider            string    `gorm:"size:100;not null;default:'Oxford University Press'" json:"provider"`                                                                                                        // Nguồn
	Language            string    `gorm:"size:20;not null;default:'es-us'" json:"language"`                                                                                                                           // Kiểu Tiếng Anh
	Tag                 string    `gorm:"size:50;not null;default:'vocabulary'" json:"tag"`                                                                                                                           // Tag, vd: vocabulary, speaking, listening
	Vi                  string    `gorm:"size:255" json:"vi"`                                                                                                                                                         // Nghĩa tiếng Việt chung
	ViNoun              string    `gorm:"size:255" json:"vi_noun"`                                                                                                                                                    // Nghĩa Tiếng Viêt theo danh từ
	ViVerb              string    `gorm:"size:255" json:"vi_verb"`                                                                                                                                                    // Nghĩa Tiếng Viêt theo động từ
	ViAdjective         string    `gorm:"size:255" json:"vi_adjective"`                                                                                                                                               // Nghĩa Tiếng Viêt theo tính từ
	ViAdverb            string    `gorm:"size:255" json:"vi_adverb"`                                                                                                                                                  // Nghĩa Tiếng Viêt theo trợ động từ
	ViPhrasal           string    `gorm:"size:255" json:"vi_phrasal"`                                                                                                                                                 // Nghĩa Tiếng Viêt theo cụm động từ
	Image               string    `gorm:"size:255;default:'https://st4.depositphotos.com/14953852/24787/v/600/depositphotos_247872612-stock-illustration-no-image-available-icon-vector.jpg'" json:"image"`           // Hình ảnh giải thích chung
	ImageNoun           string    `gorm:"size:255;default:'https://st4.depositphotos.com/14953852/24787/v/600/depositphotos_247872612-stock-illustration-no-image-available-icon-vector.jpg'" json:"image_noun"`      // Hình ảnh giải thích danh từ
	ImageVerb           string    `gorm:"size:255;default:'https://st4.depositphotos.com/14953852/24787/v/600/depositphotos_247872612-stock-illustration-no-image-available-icon-vector.jpg'" json:"image_verb"`      // Hình ảnh giải thích động từ
	ImageAdjective      string    `gorm:"size:255;default:'https://st4.depositphotos.com/14953852/24787/v/600/depositphotos_247872612-stock-illustration-no-image-available-icon-vector.jpg'" json:"image_adjective"` // Hình ảnh giải thích tính từ
	ImageAdverb         string    `gorm:"size:255;default:'https://st4.depositphotos.com/14953852/24787/v/600/depositphotos_247872612-stock-illustration-no-image-available-icon-vector.jpg'" json:"image_adverb"`    // Hình ảnh giải thích trợ động từ
	ImagePhrasel        string    `gorm:"size:255;default:'https://st4.depositphotos.com/14953852/24787/v/600/depositphotos_247872612-stock-illustration-no-image-available-icon-vector.jpg'" json:"image_phrasel"`   // Hình ảnh giải thích cụm động từ
	DefinitionNoun      string    `gorm:"size:255" json:"definition_noun"`                                                                                                                                            // Giải thích nghĩa danh từ
	DefinitionVerb      string    `gorm:"size:255" json:"definition_verb"`                                                                                                                                            // Giải thích nghĩa động từ
	DefinitionAdjective string    `gorm:"size:255" json:"definition_adjective"`                                                                                                                                       // Giải thích nghĩa tính từ
	DefinitionAdverb    string    `gorm:"size:255" json:"definition_adverb"`                                                                                                                                          // Giải thích nghĩa trợ động từ
	DefinitionPhrasal   string    `gorm:"size:255" json:"definition_phrasal"`                                                                                                                                         // Giải thích cụm động từ
	ExamplesNoun        string    `gorm:"type:text" json:"examples_noun"`                                                                                                                                             // Ví dụ về danh từ
	ExamplesVerb        string    `gorm:"type:text" json:"examples_verb"`                                                                                                                                             // Ví dụ về động từ
	ExamplesAdjective   string    `gorm:"type:text" json:"examples_adjective"`                                                                                                                                        // Ví dụ về tính từ
	ExamplesAdverb      string    `gorm:"type:text" json:"examples_adverb"`                                                                                                                                           // Ví dụ về trợ động từ
	ExamplesPhrasal     string    `gorm:"type:text" json:"examples_phrasal"`                                                                                                                                          // Ví dụ về cụm động từ
	AudioFile           string    `gorm:"size:255" json:"audio_file"`                                                                                                                                                 // Phát âm chuẩn bản xứ
	Dialects            string    `gorm:"size:100:default:'American English'" json:"dialects"`                                                                                                                        // Kiểu Tiếng Anh, vd: American English
	PhoneticSpelling    string    `gorm:"size:100" json:"phonetic_spelling"`                                                                                                                                          // Phiên âm
	LexicalCategory     string    `gorm:"size:50" json:"lexicalCategory"`                                                                                                                                             // Đóng Vai trò gì -> chung, vd: noun | verb | adjective | adverb | phrasal
	Unit                string    `gorm:"size:50;default:'Unit 1'" json:"unit"`                                                                                                                                       // Unit trong sách
	Book                string    `gorm:"size:100;default:'Destination B1'" json:"book"`                                                                                                                              // Sách
	Level               string    `gorm:"size:3;default:'A1'" json:"level"`                                                                                                                                           // Trình độ từ vựng, vd: A2, C1
	CreatedAt           time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt           time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

type CountVocabulary struct {
	Count uint32
}

type GetStudyVocab struct {
	Word string `gorm:"size:255;not null;unique" json:"word"` // Từ ngữ - Tiếng Anh
}

type CountAwlGroup struct {
	Count uint32
}
