package httpx

import (
	"encoding/json"
	"net/http"
)

type Code string

const (
	CodeInvalidID           Code = "Invalid_id"
	CodeInternalServerError Code = "Internal_server_error"
	CodeInvalidBody         Code = "Invalid_body"
	CodeValidationFailed    Code = "vaidation_failed"
)

type errorEnvelop struct {
	Error errorPayload `json:"error"`
}

type errorPayload struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, status int, message string, code Code) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorEnvelop{Error: errorPayload{Code: code, Message: message}})
}
