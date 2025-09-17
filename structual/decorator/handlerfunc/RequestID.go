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

func RequestIDChecker(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		xRequestID := r.Header.Get("X-Request-ID")
		if strings.TrimSpace(xRequestID) == "" {
			xRequestID = randomString(16)
			r.Header.Set("X-Request-ID", xRequestID)
		}
		respXRequestID := w.Header().Get("X-Request-ID")
		if strings.TrimSpace(respXRequestID) == "" {
			w.Header().Set("X-Request-ID", xRequestID)
		}

		h(w, r)
	}
}
