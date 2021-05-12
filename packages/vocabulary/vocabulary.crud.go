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
		var inputExample []string
		for _, value := range oxfordCrudData.Examples {
			inputExample = append(inputExample, fmt.Sprintf(`"%s"`, value.Text))
		}
		createData := entities.Vocabulary{
			Word:             oxfordCrudData.ID,
			Provider:         oxfordCrudData.Provider,
			Language:         oxfordCrudData.Language,
			Vi:               addVocabData.Vi,
			Image:            addVocabData.Image,
			Definition:       strings.Replace(ParseArrayToString(oxfordCrudData.Definition), "'", "", -1),
			Examples:         strings.Replace(ParseArrayToStringFor2More(inputExample), "'", "", -1),
			AudioFile:        oxfordCrudData.AudioFile,
			Dialects:         strings.Replace(ParseArrayToString(oxfordCrudData.Dialects), "'", "", -1),
			PhoneticSpelling: oxfordCrudData.PhoneticSpelling,
			Unit:             addVocabData.Unit,
			Book:             addVocabData.Book,
			Level:            addVocabData.Level,
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
	return fmt.Sprintf(`["%s"]`, strings.Join(input, ", "))
}

func ParseArrayToStringFor2More(input []string) string {
	return fmt.Sprintf("[%s]", strings.Join(input, ", "))
}

func DetechOxfordRespData(vocabRespData OxfordRespJSON) OxfordCRUDJSON {
	var oxfordCrudData OxfordCRUDJSON
	var getOxfordExamples []OxfordResultLexical2EntriesSensesExamplesJSON
	for index, results := range vocabRespData.Results {
		if index == 0 {
			for _, lexicalEntries := range results.LexicalEntries {
				for _, entries := range lexicalEntries.Entries {
					for _, senses := range entries.Senses {
						for _, text := range senses.Examples {
							getOxfordExamples = append(getOxfordExamples, text)
							oxfordCrudData = OxfordCRUDJSON{
								ID:         vocabRespData.ID,
								Provider:   vocabRespData.Metadata.Provider,
								Language:   results.Language,
								Definition: senses.Definitions,
								Examples:   getOxfordExamples,
							}
						}
					}
					for _, pronunciations := range entries.Pronunciations {
						oxfordCrudData.AudioFile = pronunciations.AudioFile
						oxfordCrudData.Dialects = pronunciations.Dialects
						oxfordCrudData.PhoneticSpelling = pronunciations.PhoneticSpelling
					}
				}
			}
		}
	}
	return oxfordCrudData
}
