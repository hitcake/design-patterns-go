package handlerfunc

import (
	"net/http"
	"testing"
)

func TestChain(t *testing.T) {
	helloHandler := &EndpointHandler{func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}}
	xRequestIDFilter := &XRquestIDHandler{}
	monitorFilter := &MonitorHandler{}
	tokenHandlerFilter := &TokenHandler{}

	xRequestIDFilter.SetNextHandler(monitorFilter)
	monitorFilter.SetNextHandler(tokenHandlerFilter)
	tokenHandlerFilter.SetNextHandler(helloHandler)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		xRequestIDFilter.Handle(w, r)
	})
	http.ListenAndServe(":8080", nil)
}
