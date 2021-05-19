package webhook

import (
	"fmt"
	"io/ioutil"
	"me-english/utils/errorcode"
	"me-english/utils/resp"
	"net/http"
)

func TelegramPushWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_400)
		return
	}
	fmt.Printf("%s\n", body)
	// textBytes := []string{}
	// err = json.Unmarshal(body, &textBytes)
	// console.Info(r)
	resp.Success(w, http.StatusOK, struct {
		Msg string `json:"msg"`
	}{
		Msg: "Thành công",
	})
	return
}
