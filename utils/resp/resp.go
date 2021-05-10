package resp

import (
	"encoding/json"
	"fmt"
	"me-english/utils/errorcode"
	"net/http"
)

type statusRespOK struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type statusRespFailed struct {
	Status int                       `json:"status"`
	Data   errorcode.ErrorCodeStruct `json:"data"`
}

/**
{
    "status": number,
    "data": []
}
*/
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	respSuccess := statusRespOK{
		Status: statusCode,
		Data:   data,
	}
	err := json.NewEncoder(w).Encode(respSuccess)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

/**
{
    "status": number,
    "data": [
        "error_code": string,
        "message": string,
    ]
}
*/
func Failed(w http.ResponseWriter, statusCode int, data errorcode.ErrorCodeStruct) {
	w.WriteHeader(statusCode)
	var respFailed = statusRespFailed{
		Status: statusCode,
		Data:   data,
	}
	err := json.NewEncoder(w).Encode(respFailed)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
