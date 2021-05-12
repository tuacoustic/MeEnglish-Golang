package vocabulary

import (
	"errors"
	"net/url"
	"strings"
)

func (p *AddVocabRequestStruct) Validation(action string) error {
	switch strings.ToLower(action) {
	case "add_vocab":
		if p.Word == "" {
			return errors.New("Chưa nhập từ mới")
		}
		if p.Vi == "" {
			return errors.New("Chưa nhập nghĩa tiếng việt")
		}
		if p.Level == "" {
			return errors.New("Chưa nhập cấp độ từ vựng")
		}
		if p.Image != "" {
			_, err := url.ParseRequestURI(p.Image)
			if err != nil {
				return errors.New("Không đúng đường dẫn ảnh")
			}
		}
		return nil
	}
	return nil
}
