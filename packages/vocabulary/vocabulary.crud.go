package vocabulary

import (
	"encoding/json"
	"errors"
	"fmt"
	"me-english/entities"
	"me-english/utils/channels"
	"me-english/utils/config"
	sendReq "me-english/utils/sendRequest"
	"strings"

	"github.com/jinzhu/gorm"
)

type repositoryVocabularyCRUD struct {
	db *gorm.DB
}

func NewRepositoryVocabularyCRUD(db *gorm.DB) *repositoryVocabularyCRUD {
	return &repositoryVocabularyCRUD{db}
}

var (
	msg = ""
)

func (r *repositoryVocabularyCRUD) AddVocab(addVocabData AddVocabRequestStruct) (bool, error) {
	var err error
	// Request đến Oxford
	// Detech
	urlSendRequest := strings.Replace(config.OXFORD_URL_API, "word_params", fmt.Sprintf("%s", addVocabData.Word), -1)
	responseData := sendReq.PostRequestToOxford(urlSendRequest, "GET", "")
	var vocabRespData OxfordRespJSON
	json.Unmarshal([]byte(responseData), &vocabRespData)
	oxfordCrudData := DetechOxfordRespData(vocabRespData)
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		var inputExamplesNoun []string
		var inputExamplesVerb []string
		var inputExamplesAdjective []string
		var inputExamplesAdverb []string
		var inputExamplesPhrasal []string

		var inputLexicalCategory []string
		// Tìm từ vựng
		countExistWord := entities.CountVocabulary{}
		queryDBByWord := fmt.Sprintf(`
select count(id) as count
from vocabulary
where word = "%s"
		`, oxfordCrudData.ID)
		r.db.Raw(queryDBByWord).Find(&countExistWord)
		if countExistWord.Count > 0 {
			msg = "Từ vựng này đã tồn tại"
			ch <- false
			return
		}
		for _, examplesNoun := range oxfordCrudData.ExamplesNoun {
			inputExamplesNoun = append(inputExamplesNoun, fmt.Sprintf(`"%s"`, examplesNoun.Text))
		}
		for _, examplesVerb := range oxfordCrudData.ExamplesVerb {
			inputExamplesVerb = append(inputExamplesVerb, fmt.Sprintf(`"%s"`, examplesVerb.Text))
		}
		for _, examplesAdjective := range oxfordCrudData.ExamplesAdjective {
			inputExamplesAdjective = append(inputExamplesAdjective, fmt.Sprintf(`"%s"`, examplesAdjective.Text))
		}
		for _, examplesAdverb := range oxfordCrudData.ExamplesAdverb {
			inputExamplesAdverb = append(inputExamplesAdverb, fmt.Sprintf(`"%s"`, examplesAdverb.Text))
		}
		for _, examplesPhrasal := range oxfordCrudData.ExamplesPhrasal {
			inputExamplesPhrasal = append(inputExamplesPhrasal, fmt.Sprintf(`"%s"`, examplesPhrasal.Text))
		}
		for _, lexicalCategory := range oxfordCrudData.LexicalCategory {
			inputLexicalCategory = append(inputLexicalCategory, fmt.Sprintf(`"%s"`, lexicalCategory))
		}
		createData := entities.Vocabulary{
			Word:                oxfordCrudData.ID,
			Provider:            oxfordCrudData.Provider,
			Language:            oxfordCrudData.Language,
			Vi:                  addVocabData.Vi,
			ViVerb:              addVocabData.ViVerb,
			ViAdjective:         addVocabData.ViAdjective,
			ViAdverb:            addVocabData.ViAdverb,
			ViPhrasal:           addVocabData.ViPhrasel,
			Image:               addVocabData.Image,
			ImageNoun:           addVocabData.ImageNoun,
			ImageVerb:           addVocabData.ImageAdverb,
			ImageAdjective:      addVocabData.ImageAdjective,
			ImageAdverb:         addVocabData.ImageAdverb,
			ImagePhrasel:        addVocabData.ImagePhrasel,
			DefinitionNoun:      strings.Replace(ParseArrayToString(oxfordCrudData.DefinitionNoun), "'", "", -1),
			DefinitionVerb:      strings.Replace(ParseArrayToString(oxfordCrudData.DefinitionVerb), "'", "", -1),
			DefinitionAdjective: strings.Replace(ParseArrayToString(oxfordCrudData.DefinitionAdjective), "'", "", -1),
			DefinitionAdverb:    strings.Replace(ParseArrayToString(oxfordCrudData.DefinitionAdverb), "'", "", -1),
			DefinitionPhrasal:   strings.Replace(ParseArrayToString(oxfordCrudData.DefinitionPhrasal), "'", "", -1),
			ExamplesNoun:        strings.Replace(ParseArrayToStringFor2More(inputExamplesNoun), "'", "", -1),
			ExamplesVerb:        strings.Replace(ParseArrayToStringFor2More(inputExamplesVerb), "'", "", -1),
			ExamplesAdjective:   strings.Replace(ParseArrayToStringFor2More(inputExamplesAdjective), "'", "", -1),
			ExamplesAdverb:      strings.Replace(ParseArrayToStringFor2More(inputExamplesAdverb), "'", "", -1),
			ExamplesPhrasal:     strings.Replace(ParseArrayToStringFor2More(inputExamplesPhrasal), "'", "", -1),
			AudioFile:           oxfordCrudData.AudioFile,
			Dialects:            strings.Replace(ParseArrayToString(oxfordCrudData.Dialects), "'", "", -1),
			PhoneticSpelling:    oxfordCrudData.PhoneticSpelling,
			Unit:                addVocabData.Unit,
			Book:                addVocabData.Book,
			Level:               addVocabData.Level,
			LexicalCategory:     strings.Replace(ParseArrayToStringFor2More(inputLexicalCategory), "'", "", -1),
			AwlGroupID:          addVocabData.AwlGroupID,
		}
		err = r.db.Debug().Model(&entities.Vocabulary{}).Create(&createData).Error
		if err != nil {
			msg = fmt.Sprintf("%s", err)
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return true, nil
	}
	return false, errors.New(msg)
}

func ParseArrayToString(input []string) string {
	if len(input) == 0 {
		return ""
	}
	return fmt.Sprintf(`["%s"]`, strings.Join(input, ", "))
}

func ParseArrayToStringFor2More(input []string) string {
	if len(input) == 0 {
		return ""
	}
	return fmt.Sprintf("[%s]", strings.Join(input, ", "))
}

func DetechOxfordRespData(vocabRespData OxfordRespJSON) OxfordCRUDJSON {
	var oxfordCrudData OxfordCRUDJSON
	var getLexicalCategory []string
	for index, results := range vocabRespData.Results {
		if index == 0 {
			for _, lexicalEntries := range results.LexicalEntries {
				switch lexicalEntries.LexicalCategory.ID {
				case NOUN:
					var getOxfordExamplesNoun []OxfordResultLexical2EntriesSensesExamplesJSON
					for _, entries := range lexicalEntries.Entries {
						for _, senses := range entries.Senses {
							for _, text := range senses.Examples {
								getOxfordExamplesNoun = append(getOxfordExamplesNoun, text)
								oxfordCrudData = OxfordCRUDJSON{
									ID:             vocabRespData.ID,
									Provider:       vocabRespData.Metadata.Provider,
									Language:       results.Language,
									DefinitionNoun: senses.Definitions,
									ExamplesNoun:   getOxfordExamplesNoun,
								}
							}
						}
						for _, pronunciations := range entries.Pronunciations {
							oxfordCrudData.AudioFile = pronunciations.AudioFile
							oxfordCrudData.Dialects = pronunciations.Dialects
							oxfordCrudData.PhoneticSpelling = pronunciations.PhoneticSpelling
						}
					}
					break
				case VERB:
					var getOxfordExamplesVerb []OxfordResultLexical2EntriesSensesExamplesJSON
					for _, entries := range lexicalEntries.Entries {
						for _, senses := range entries.Senses {
							for _, text := range senses.Examples {
								getOxfordExamplesVerb = append(getOxfordExamplesVerb, text)
								oxfordCrudData = OxfordCRUDJSON{
									ID:             vocabRespData.ID,
									Provider:       vocabRespData.Metadata.Provider,
									Language:       results.Language,
									DefinitionVerb: senses.Definitions,
									ExamplesVerb:   getOxfordExamplesVerb,
								}
							}
						}
						for _, pronunciations := range entries.Pronunciations {
							oxfordCrudData.AudioFile = pronunciations.AudioFile
							oxfordCrudData.Dialects = pronunciations.Dialects
							oxfordCrudData.PhoneticSpelling = pronunciations.PhoneticSpelling
						}
					}
					break
				case ADJECTIVE:
					var getOxfordExamplesAdjective []OxfordResultLexical2EntriesSensesExamplesJSON
					for _, entries := range lexicalEntries.Entries {
						for _, senses := range entries.Senses {
							for _, text := range senses.Examples {
								getOxfordExamplesAdjective = append(getOxfordExamplesAdjective, text)
								oxfordCrudData = OxfordCRUDJSON{
									ID:                  vocabRespData.ID,
									Provider:            vocabRespData.Metadata.Provider,
									Language:            results.Language,
									DefinitionAdjective: senses.Definitions,
									ExamplesAdjective:   getOxfordExamplesAdjective,
								}
							}
						}
						for _, pronunciations := range entries.Pronunciations {
							oxfordCrudData.AudioFile = pronunciations.AudioFile
							oxfordCrudData.Dialects = pronunciations.Dialects
							oxfordCrudData.PhoneticSpelling = pronunciations.PhoneticSpelling
						}
					}
					break
				case ADVERB:
					var getOxfordExamplesAdverb []OxfordResultLexical2EntriesSensesExamplesJSON
					for _, entries := range lexicalEntries.Entries {
						for _, senses := range entries.Senses {
							for _, text := range senses.Examples {
								getOxfordExamplesAdverb = append(getOxfordExamplesAdverb, text)
								oxfordCrudData = OxfordCRUDJSON{
									ID:               vocabRespData.ID,
									Provider:         vocabRespData.Metadata.Provider,
									Language:         results.Language,
									DefinitionAdverb: senses.Definitions,
									ExamplesAdverb:   getOxfordExamplesAdverb,
								}
							}
						}
						for _, pronunciations := range entries.Pronunciations {
							oxfordCrudData.AudioFile = pronunciations.AudioFile
							oxfordCrudData.Dialects = pronunciations.Dialects
							oxfordCrudData.PhoneticSpelling = pronunciations.PhoneticSpelling
						}
					}
					break
				case PHRASAL:
					var getOxfordExamplesPhrasal []OxfordResultLexical2EntriesSensesExamplesJSON
					for _, entries := range lexicalEntries.Entries {
						for _, senses := range entries.Senses {
							for _, text := range senses.Examples {
								getOxfordExamplesPhrasal = append(getOxfordExamplesPhrasal, text)
								oxfordCrudData = OxfordCRUDJSON{
									ID:                vocabRespData.ID,
									Provider:          vocabRespData.Metadata.Provider,
									Language:          results.Language,
									DefinitionPhrasal: senses.Definitions,
									ExamplesPhrasal:   getOxfordExamplesPhrasal,
								}
							}
						}
						for _, pronunciations := range entries.Pronunciations {
							oxfordCrudData.AudioFile = pronunciations.AudioFile
							oxfordCrudData.Dialects = pronunciations.Dialects
							oxfordCrudData.PhoneticSpelling = pronunciations.PhoneticSpelling
						}
					}
					break
				}
				getLexicalCategory = append(getLexicalCategory, lexicalEntries.LexicalCategory.ID)
				oxfordCrudData.LexicalCategory = getLexicalCategory
			}

		}
	}
	return oxfordCrudData
}
