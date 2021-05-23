package main

import (
	"fmt"
	"log"
	"me-english/database"
	"me-english/packages/router"
	"me-english/utils/config"
	"me-english/utils/console"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	tb "gopkg.in/tucnak/telebot.v2"
)

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port),
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))
}

func main() {
	// config.Load()
	// Auto Generate Model Schema
	status := database.Auto()
	if status == true {
		console.Info("Connect Mysql Successful")
	}
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "http://203.162.54.20:8888",

		Token:  config.TELEGRAM_TOKEN_MEENGLISH,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	b.Start()
	console.Info("Listening [::]:", config.PORT)
	listen(config.PORT)
}
