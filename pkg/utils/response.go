package utils

import (
	"net/http"
)

func Success(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func Fail(w http.ResponseWriter, status int, message []byte) {
	w.WriteHeader(status)
	w.Write(message)
}
