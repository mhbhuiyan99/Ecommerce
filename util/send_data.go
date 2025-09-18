package util

import (
	"net/http"
	"encoding/json"
)

func SendData(w http.ResponseWriter, statusCode int, data interface{}) { // data interface{} for getting any types of data
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w) 
	encoder.Encode(data) 
}

func SendError(w http.ResponseWriter, statusCode int, msg string){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}