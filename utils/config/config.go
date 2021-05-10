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
	MYSQL_DB_URL           = "tuacoustic:Tudeptrai0312@tcp(192.168.170.86:3306)/me-english?charset=utf8&parseTime=True&loc=Local"
	MYSQL_DB_IP            = ""
	MYSQL_DB_DRIVER        = "mysql"
	REDIS_ADDR             = "localhost:6379"
	REDIS_PASS             = ""
	SECRETKEY       []byte = []byte("66f3cca610bad24b27857bbc4695dbeb")
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
