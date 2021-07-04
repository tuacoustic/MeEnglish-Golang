package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT = 4040
	// MYSQL_DB_URL           = "gkitchen:6CUykhj45tKgBJcs@tcp(10.0.1.233:3306)/gkitchen?charset=utf8&parseTime=True&loc=Local"
	MYSQL_DB_URL           = "tuacoustic:Curveruler0312@tcp(192.168.170.86:3306)/me_english?charset=utf8&parseTime=True&loc=Local"
	MYSQL_DB_IP            = ""
	MYSQL_DB_DRIVER        = "mysql"
	REDIS_ADDR             = "localhost:6379"
	REDIS_PASS             = ""
	SECRETKEY       []byte = []byte("66f3cca610bad24b27857bbc4695dbeb")

	// Oxford Config - 1k req/month
	OXFORD_APP_ID  = "9142aa26"
	OXFORD_APP_KEY = "8c069747485a726791b0f68f6829af10"
	OXFORD_URL_API = "https://od-api.oxforddictionaries.com/api/v2/entries/en-us/word_params?strictMatch=false"

	// Telegram config - Limited 20reqs/s
	TELEGRAM_SEND_MESSAGE    = "https://api.telegram.org/bot@token_params/sendMessage?@params"
	TELEGRAM_TOKEN_MEENGLISH = "1824373162:AAHrLY0caFSNJVaZI17B7pPxzQ_dw73YRBU"
)

func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	MYSQL_DB_DRIVER = os.Getenv("MYSQL_DB_DRIVER")
	MYSQL_DB_URL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_DB_USER"),
		os.Getenv("MYSQL_DB_PASS"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DB_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	)
	REDIS_ADDR = fmt.Sprintf("%s:%s", os.Getenv("REDIS_IP"), os.Getenv("REDIS_PORT"))
	REDIS_PASS = os.Getenv("REDIS_PASS")
}
