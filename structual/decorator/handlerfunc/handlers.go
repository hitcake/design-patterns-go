package handlerfunc

import "net/http"

func NeedAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func NoNeedAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
