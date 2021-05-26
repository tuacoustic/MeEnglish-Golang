package main

import (
	"fmt"
	"log"
	"me-english/database"
	"me-english/packages/router"
	"me-english/packages/webhook"
	"me-english/utils/config"
	"me-english/utils/console"
	"net/http"

	"github.com/gorilla/handlers"
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
	webhook.ConnectWebhook()
	console.Info("Listening [::]:", config.PORT)
	listen(config.PORT)
}
