package handlerfunc

import "net/http"

type EndpointHandler struct {
	serv http.HandlerFunc
}

func (h *EndpointHandler) Handle(w http.ResponseWriter, r *http.Request) {
	h.serv(w, r)
}
func (h *EndpointHandler) SetNextHandler(handler Handler) {

}
