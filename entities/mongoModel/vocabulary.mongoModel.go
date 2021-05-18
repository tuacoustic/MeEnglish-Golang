package mongoModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vocabulary struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Word             string             `json:"word" bson:"word"`
	Provider         string             `json:"provider" bson:"provider"`
	Language         string             `json:"language" bson:"language"`
	Tag              string             `json:"tag" bson:"tag"`
	Vi               string             `json:"vi" bson:"vi"`
	Image            string             `json:"image" bson:"image"`
	Definition       string             `json:"definition" bson:"definition"`
	Examples         string             `json:"example" bson:"example"`
	AudioFile        string             `json:"audio_file" bson:"audio_file"`
	Dialects         string             `json:"dialects" bson:"dialects"`
	PhoneticSpelling string             `json:"phonetic_spelling" bson:"phonetic_spelling"`
	Unit             string             `json:"unit" bson:"unit"`
	Book             string             `json:"book" bson:"book"`
	Level            string             `json:"level" bson:"level"`
	LexicalCategory  string             `json:"lexicalCategory" bson:"lexicalCategory"`
}
