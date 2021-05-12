package vocabulary

import (
	"encoding/json"
	"io/ioutil"
	"me-english/database"
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
	// Validation
	err = addVocabData.Validation("add_vocab")
	if err != nil {
		respErr := errorcode.CustomErr(ERROR_400, err)
		resp.Failed(w, http.StatusBadRequest, respErr)
		return
	}
	db, err := database.MysqlConnect()
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_500)
		return
	}
	repo := NewRepositoryVocabularyCRUD(db)
	func(vocabularyRepo VocabularyRepository) {
		status, err := vocabularyRepo.AddVocab(addVocabData)
		if err != nil {
			respErr := errorcode.CustomErr(ERROR_400, err)
			resp.Failed(w, http.StatusBadRequest, respErr)
		}
		if status == true {
			resp.Success(w, http.StatusOK, struct {
				Msg string `json:"msg"`
			}{
				Msg: "Thêm từ mới thành công",
			})
		}
		return
	}(repo)
}
