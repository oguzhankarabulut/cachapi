package api

import (
	"fmt"
	"log"
	"net/http"
)

// LogRequest log url, method and header for all requests. It works as a middleware.
func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("url: %s method: %s, header: %s", r.URL, r.Method, r.Header))
		h.ServeHTTP(w, r)
	}
}
