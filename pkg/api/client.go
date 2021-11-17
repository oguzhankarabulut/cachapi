package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	errMethodNotAllowed = errors.New("method not allowed")
)

type Response struct {
	Content      interface{} `json:"content"`
	Message      string      `json:"message"`
	ResponseCode int         `json:"responseCode"`
}

func newResponse(
	content interface{},
	message string,
	responseCode int,
) *Response {
	return &Response{
		Content:      content,
		Message:      message,
		ResponseCode: responseCode,
	}
}

// write return response
func write(w http.ResponseWriter, r *Response) {
	o, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(o)
}

// writeOK should be used for status 200 responses
func writeOK(w http.ResponseWriter, v interface{}) {
	r := newResponse(v, "", http.StatusOK)
	write(w, r)
}

// writeBadRequest should be used for status 400 responses
func writeBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	res := newResponse(nil, err.Error(), http.StatusBadRequest)
	w.WriteHeader(http.StatusBadRequest)
	write(w, res)
}

// writeNotFound should be used for 404 responses
func writeNotFound(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	res := newResponse(nil, err.Error(), http.StatusNotFound)
	w.WriteHeader(http.StatusNotFound)
	write(w, res)
}

// writeMethodErr should be used for 405 responses
func writeMethodErr(w http.ResponseWriter, r *http.Request) {
	logError(r, errMethodNotAllowed)
	res := newResponse(nil, errMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	w.WriteHeader(http.StatusMethodNotAllowed)
	write(w, res)
}

// logError log errors if response is not success
func logError(r *http.Request, err error) {
	log.Println(fmt.Sprintf("Error: %s Url: %s, Method: %s", err.Error(), r.URL, r.Method))
}
