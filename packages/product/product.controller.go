package product

import (
	"me-english/utils/resp"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	resp.Success(w, http.StatusOK, struct {
		Msg string `json:"msg"`
	}{
		Msg: "Xin chào",
	})
}
