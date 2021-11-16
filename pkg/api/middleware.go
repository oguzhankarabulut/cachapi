package api

import (
	"fmt"
	"log"
	"net/http"
)

func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("url: %s", r.URL))
		h.ServeHTTP(w, r)
	}
}
