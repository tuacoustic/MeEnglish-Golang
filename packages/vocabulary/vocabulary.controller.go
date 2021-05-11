package vocabulary

import (
	"encoding/json"
	"io/ioutil"
	"me-english/database"
	"me-english/utils/console"
	"me-english/utils/errorcode"
	"me-english/utils/resp"
	"net/http"
)

func AddVocab(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Data from Request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_400)
		return
	}
	var addVocabData AddVocabRequestStruct
	err = json.Unmarshal(body, &addVocabData)
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_400)
		return
	}
	// Detech
	var vocabRespData OxfordRespJSON
	err = json.Unmarshal([]byte(OxfordResp), &vocabRespData)
	if err != nil {
		console.Info(err)
		resp.Failed(w, http.StatusBadRequest, errorcode.CustomErr(ERROR_400, err))
	}
	oxfordCrudData := DetechOxfordRespData(vocabRespData)
	db, err := database.MysqlConnect()
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_500)
		return
	}
	repo := NewRepositoryVocabularyCRUD(db)
	func(vocabularyRepo VocabularyRepository) {
		status, err := vocabularyRepo.AddVocab(oxfordCrudData)
		if err != nil {
			respErr := errorcode.CustomErr(ERROR_400, err)
			resp.Failed(w, http.StatusBadRequest, respErr)
		}
		console.Info(addVocabData.Word)
		if status == true {
			resp.Success(w, http.StatusOK, struct {
				Msg string `json:"msg"`
			}{
				Msg: "Xin chào",
			})
		}
		return
	}(repo)
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
