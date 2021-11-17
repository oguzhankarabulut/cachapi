package api

import "net/http"

type cacheRequest struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

// key gets key's value from query parameter
func key(r *http.Request) string {
	return r.URL.Query().Get("key")
}
