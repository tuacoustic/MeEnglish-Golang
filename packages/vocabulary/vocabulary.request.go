package vocabulary

type AddVocabRequestStruct struct {
	Word           string `json:"word"`
	Vi             string `json:"vi"`
	Level          string `json:"level"`
	Book           string `json:"book"`
	Image          string `json:"image"`
	Unit           string `json:"unit"`
	ViNoun         string `json:"vi_noun"`
	ViVerb         string `json:"vi_verb"`
	ViAdjective    string `json:"vi_adjective"`
	ViAdverb       string `json:"vi_adverb"`
	ViPhrasel      string `json:"vi_phrasel"`
	ImageNoun      string `json:"image_noun"`
	ImageVerb      string `json:"image_verb"`
	ImageAdjective string `json:"image_adjective"`
	ImageAdverb    string `json:"image_adverb"`
	ImagePhrasel   string `json:"image_phrasel"`
	AwlGroupID     uint64 `json:"awl_group_id"`
}
