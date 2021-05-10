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
	var addVocabData AddVocabReuqestStruct
	err = json.Unmarshal(body, &addVocabData)
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_400)
		return
	}
	db, err := database.MysqlConnect()
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_500)
		return
	}
	repo := NewRepositoryVocabularyCRUD(db)
	func(vocabulary VocabularyRepository) {

	}(repo)
	console.Info(addVocabData.Word)
	resp.Success(w, http.StatusOK, struct {
		Msg string `json:"msg"`
	}{
		Msg: "Xin chào",
	})
	return
}
