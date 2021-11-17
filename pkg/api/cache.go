package api

import (
	"cachapi/pkg/facade"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errKeyNonExist = errors.New("key is not exist at request")
	errKeyNil = errors.New("key can not be nil")
	errValueNil = errors.New("value can not be nil")
)

type CacheHandler struct {
	 cacheFacade facade.CacheFacade
}

func NewCacheHandler(cf facade.CacheFacade) *CacheHandler {
	return &CacheHandler{
		cacheFacade: cf,
	}
}

type GetRequest struct {
	Key string `json:"key"`
}

type SetRequest struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

// HandleCache controls request method and execute necessary method or return method error
func (c CacheHandler) HandleCache(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodGet:
		c.Get(w, r)
	case http.MethodPost:
		c.Set(w, r)
	default:
		writeMethodErr(w, r)
	}
}

// HandleFlush control method method
func (c CacheHandler) HandleFlush(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.FlushCache(w, r)
	default:
		writeMethodErr(w, r)
	}
}

func (c CacheHandler) Get(w http.ResponseWriter, r *http.Request) {
	k := key(r)
	if len(k) == 0 {
		writeBadRequest(w, r, errKeyNonExist)
		return
	}
	nc, err :=c.cacheFacade.Get(k)
	if err != nil {
		writeNotFound(w, r, err)
		return
	}
	writeOK(w, NewSingleResponse(nc.GetKey(), nc.GetValue()))
}

func (c CacheHandler) Set(w http.ResponseWriter, r *http.Request) {
	cr := new(cacheRequest)
	if err := json.NewDecoder(r.Body).Decode(cr); err != nil {
		writeBadRequest(w, r, err)
		return
	}
	if cr.Key == "" {
		writeBadRequest(w, r, errKeyNil)
		return
	}
	if cr.Value == nil {
		writeBadRequest(w, r, errValueNil)
		return
	}
	nc := c.cacheFacade.Set(cr.Key, cr.Value)

	writeOK(w, NewSingleResponse(nc.GetKey(), nc.GetValue()))
}

func (c CacheHandler) FlushCache(w http.ResponseWriter, r *http.Request) {
	c.cacheFacade.FlushCache()
	writeOK(w, nil)
}
