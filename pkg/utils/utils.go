package utils

import (
	"encoding/json"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Original-Forwarded-For")
	if IPAddress != "" {
		return IPAddress
	}

	IPAddress = r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func ValidateAndProcessData(w http.ResponseWriter, r *http.Request, data interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return false
	}

	if err := validate.Struct(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	return true
}
