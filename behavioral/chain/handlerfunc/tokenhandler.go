package handlerfunc

import (
	"net/http"
	"strings"
)

type TokenHandler struct {
	AbstractHandler
}

func (h *TokenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		if h.nextHandler != nil {
			h.nextHandler.Handle(w, r)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
