package api

import (
	"cachapi/pkg/facade"
	"net/http"
)

const address = ":3000"

func Server() {
	cf := facade.NewCacheFacade()
	handler := NewCacheHandler(cf)
	http.HandleFunc("/api/v1/cache", LogRequest(handler.HandleCache))
	http.HandleFunc("/api/v1/flush", LogRequest(handler.HandleFlush))
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
