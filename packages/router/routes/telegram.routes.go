package routes

import (
	webhook "me-english/packages/webhook"
	"net/http"
)

var telegramRoutes = []Route{
	{
		Uri:          "/telegram/me-english",
		Method:       http.MethodPost,
		Handler:      webhook.TelegramPushWebhook,
		CheckHeaders: false,
	},
}
