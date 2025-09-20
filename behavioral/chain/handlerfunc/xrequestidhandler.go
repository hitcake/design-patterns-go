package handlerfunc

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type XRquestIDHandler struct {
	AbstractHandler
}

func (h *XRquestIDHandler) Handle(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-Id")
	if requestID == "" {
		xRequestID := randomString(16)
		r.Header.Set("X-Request-Id", xRequestID)
	}
	respXRequestID := w.Header().Get("X-Request-ID")
	if strings.TrimSpace(respXRequestID) == "" {
		w.Header().Set("X-Request-Id", requestID)
	}
	if h.nextHandler != nil {
		h.nextHandler.Handle(w, r)
	}
}
