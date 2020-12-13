package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type defaultError struct {
	StatusCode int       `json:"code"`
	Message    string    `json:"mensagem"`
	Error      string    `json:"erro"`
	TimeError  time.Time `json:"timestamp"`
	Path       string    `json:"path"`
}

var httpStatusMessage = map[int]string{
	400: "Requisição com campos inválidos",
	404: "Produto não encontrado",
	500: "Um erro ocorreu e não foi possível processar a requisição",
}

// HandlerError - Global HandlerError function
func HandlerError(w http.ResponseWriter, r *http.Request, status int, message string) {
	de := defaultError{
		Message:    httpStatusMessage[status],
		StatusCode: status,
		TimeError:  time.Now(),
		Error:      message,
		Path:       r.RequestURI,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(de.StatusCode)
	_ = json.NewEncoder(w).Encode(&de)
}
