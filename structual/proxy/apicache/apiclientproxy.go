package apicache

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	resp     []byte
	err      error
	expireAt time.Time
}
type ApiClientProxy struct {
	apiClient     ApiClient
	cache         map[string]result
	cacheLock     sync.RWMutex
	expireSeconds int64
}

func NewApiClientProxy(apiClient ApiClient, expireSeconds int64) *ApiClientProxy {
	return &ApiClientProxy{apiClient: apiClient, cache: make(map[string]result), cacheLock: sync.RWMutex{}, expireSeconds: expireSeconds}
}

func (a *ApiClientProxy) Do(req *http.Request) ([]byte, error) {
	if req.Method != "GET" {
		return a.apiClient.Do(req)
	} else {
		a.cacheLock.RLock()
		if response, ok := a.cache[req.URL.String()]; ok {
			if response.expireAt.After(time.Now()) {
				a.cacheLock.RUnlock()
				return response.resp, response.err
			}
			fmt.Printf("cache expired for %s\n", req.URL.String())
		}
		a.cacheLock.RUnlock()

		a.cacheLock.Lock()
		defer a.cacheLock.Unlock()
		if response, ok := a.cache[req.URL.String()]; ok {
			if response.expireAt.After(time.Now()) {
				return response.resp, response.err
			} else {
				delete(a.cache, req.URL.String())
			}
		}
		response, err := a.apiClient.Do(req)
		a.cache[req.URL.String()] = result{resp: response, err: err, expireAt: time.Now().Add(time.Duration(a.expireSeconds) * time.Second)}
		return response, err
	}
}
