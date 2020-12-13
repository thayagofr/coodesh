package utils

import (
	"log"
	"net/http"
	"os"
)

func LogMiddleware(next http.Handler) http.Handler {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			infoLog.Printf(" %s - %s %s %s ", r.RemoteAddr, r.Proto, r.Method, r.URL)
			next.ServeHTTP(w,r)
		})
}
