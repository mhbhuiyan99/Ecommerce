package util

import (
	"net/http"
	"encoding/json"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) { // data interface{} for getting any types of data
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w) 
	encoder.Encode(data) 
}

func SendError(w http.ResponseWriter, statusCode int, msg string){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}