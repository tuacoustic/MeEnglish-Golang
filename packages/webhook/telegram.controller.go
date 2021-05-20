package webhook

import (
	"encoding/json"
	"io/ioutil"
	"me-english/database"
	"me-english/utils/errorcode"
	"me-english/utils/resp"
	"net/http"
)

func TelegramPushWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var commandFlag bool = false
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_400)
		return
	}
	// fmt.Printf("%s\n", body)
	var telegramPushWB TelegramRespJSON
	err = json.Unmarshal([]byte(body), &telegramPushWB)
	// Người dùng /start đăng ký học
	for _, entityType := range telegramPushWB.Message.Entities {
		if entityType.Type == BOT_COMMAND {
			commandFlag = true
		}
	}
	db, err := database.MysqlConnect()
	if err != nil {
		resp.Failed(w, http.StatusBadRequest, errorcode.GeneralErr.ERR_500)
		return
	}
	repo := NewRepositoryTelegramCRUD(db)
	switch commandFlag {
	case true: // Dùng lệnh
		// Xử lý khi khách nhập start -> Tạo User vào Database
		func(telegramRepo TelegramRepository) {
			status, err := telegramRepo.CreateUser(telegramPushWB)
			if err != nil {
				respErr := errorcode.CustomErr(ERROR_400, err)
				resp.Failed(w, http.StatusBadRequest, respErr)
			}
			if status == true {
				resp.Success(w, http.StatusOK, struct {
					Msg string `json:"msg"`
				}{
					Msg: "Thêm người dùng thành công",
				})
			}
			return
		}(repo)
		break
	case false: // Không dùng lệnh

		break
	}
	return
}
