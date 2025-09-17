package handlerfunc

import (
	"net/http"
	"testing"
)

func TestDecorator(t *testing.T) {
	http.HandleFunc("/", RequestIDChecker(MonitorHandler(NoNeedAuthHandler)))
	http.HandleFunc("/needauth", RequestIDChecker(MonitorHandler(AuthHandler(NeedAuthHandler))))
	http.ListenAndServe(":8080", nil)
}
