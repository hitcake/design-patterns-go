package handlerfunc

import "net/http"

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
	SetNextHandler(handler Handler)
}

type AbstractHandler struct {
	nextHandler Handler
}

func (a *AbstractHandler) SetNextHandler(handler Handler) {
	a.nextHandler = handler
}
