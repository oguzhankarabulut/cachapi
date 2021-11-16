package api

import "net/http"

type cacheRequest struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

func key(r *http.Request) string {
	return r.URL.Query().Get("key")
}
