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
		if p.ImageAdjective != "" {
			_, err := url.ParseRequestURI(p.ImageAdjective)
			if err != nil {
				return errors.New("ImageAdjective -> Không đúng đường dẫn ảnh")
			}
		}
		if p.ImageAdverb != "" {
			_, err := url.ParseRequestURI(p.ImageAdverb)
			if err != nil {
				return errors.New("image_adverb -> Không đúng đường dẫn ảnh")
			}
		}
		if p.ImageNoun != "" {
			_, err := url.ParseRequestURI(p.ImageNoun)
			if err != nil {
				return errors.New("image_noun -> Không đúng đường dẫn ảnh")
			}
		}
		if p.ImagePhrasel != "" {
			_, err := url.ParseRequestURI(p.ImagePhrasel)
			if err != nil {
				return errors.New("image_phrasel -> Không đúng đường dẫn ảnh")
			}
		}
		if p.ImageVerb != "" {
			_, err := url.ParseRequestURI(p.ImageVerb)
			if err != nil {
				return errors.New("image_verb -> Không đúng đường dẫn ảnh")
			}
		}
		return nil
	}
	return nil
}
