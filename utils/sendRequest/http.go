package sendReq

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"me-english/utils/config"
	"me-english/utils/console"
	"net/http"
)

type requestHeaderStruct struct {
	Key   string
	Value string
}

var (
	requestHeader = []requestHeaderStruct{
		{
			Key:   "Content-Type",
			Value: "application/json",
		},
		{
			Key:   "app_id",
			Value: config.OXFORD_APP_ID,
		},
		{
			Key:   "app_key",
			Value: config.OXFORD_APP_KEY,
		},
	}
)

func PostRequestToOxford(url string, method string, jsonData interface{}) string {
	bufData, err := json.Marshal(jsonData)
	if err != nil {
		console.Info("http.go PostRequestToOxford err: %s", err)
		return ""
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bufData))
	if err != nil {
		console.Info("http.go PostRequestToOxford req err: %s", err)
		return ""
	}
	for _, sendHeaderData := range requestHeader {
		req.Header.Set(sendHeaderData.Key, sendHeaderData.Value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func PostRequestToTelegram(url string, method string, jsonData interface{}) string {
	console.Info(url)
	bufData, err := json.Marshal(jsonData)
	if err != nil {
		console.Info("http.go PostRequestToTelegram err: %s", err)
		return ""
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bufData))
	if err != nil {
		console.Info("http.go PostRequestToTelegram req err: %s", err)
		return ""
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
