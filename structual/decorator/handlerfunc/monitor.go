package handlerfunc

import (
	"fmt"
	"net/http"
	"time"
)

func MonitorHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		duration := time.Since(start)
		fmt.Printf("MonitorHandler %s took %s\n", r.URL.Path, duration)
	}
}
