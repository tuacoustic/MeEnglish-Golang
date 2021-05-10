package cache

import (
	"encoding/json"
	"me-english/entities"
	"time"
)

var (
	// Table_name:GO_Name
	SET_VOCABULARY = "Vocabulary:GO_VOCAB/vocab?id="
)

type VocabCacheStruct struct {
	Data entities.Vocabulary `json:"data"`
}

func SetVocab(key string, value VocabCacheStruct, expires time.Duration) {
	client := redisConnect()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key, json, expires*time.Second)
}

func GetVocab(key string) VocabCacheStruct {
	client := redisConnect()

	val, err := client.Get(key).Result()
	if err != nil {
		return VocabCacheStruct{}
	}

	product := VocabCacheStruct{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		panic(err)
	}
	return product
}
