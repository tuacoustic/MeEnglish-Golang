package webhook

type TelegramRespJSON struct {
	Message TelegramRespMessageJSON `json:"message"`
}

type TelegramErrJSON struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int32  `json:"error_code"`
	Description string `json:"description"`
}
type TelegramRespMessageJSON struct {
	MessageID uint64                      `json:"message_id"`
	Chat      TelegramRespChatJSON        `json:"chat"`
	Text      string                      `json:"text"`
	Entities  []TelegramRespEntitiesJSON  `json:"entities"`
	From      TelegramRespMessageFromJSON `json:"from"`
	Date      string                      `json:"date"`
}

type TelegramRespChatJSON struct {
	Type string `json:"type"`
}

type TelegramRespEntitiesJSON struct {
	Type string `json:"type"`
}

type TelegramRespMessageFromJSON struct {
	ID        uint64 `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}
