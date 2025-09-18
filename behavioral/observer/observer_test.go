package observer

import (
	"testing"
	"time"
)

func TestNotify(t *testing.T) {
	consumer01 := NewNativeListener("consumer01")
	consumer02 := NewRemoteListener("consumer02", "http://192.168.1.23:8080")
	center := NewCenter()
	center.Register("serviceA", consumer01)
	center.Register("serviceA", consumer02)
	center.Up("serviceA", "http://192.168.1.55:8080")
	center.Up("serviceA", "http://192.168.1.188:8080")
	time.Sleep(1 * time.Second)
	center.Down("serviceA", "http://192.168.1.188:8080")
	time.Sleep(1 * time.Second)
}
