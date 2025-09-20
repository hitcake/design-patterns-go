package handlerfunc

import (
	"fmt"
	"net/http"
	"time"
)

type MonitorHandler struct {
	AbstractHandler
}

func (h *MonitorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if h.nextHandler != nil {
		h.nextHandler.Handle(w, r)
	}
	elapsed := time.Since(start)
	fmt.Printf(" url handle with elapsed %s\n", elapsed)
}
