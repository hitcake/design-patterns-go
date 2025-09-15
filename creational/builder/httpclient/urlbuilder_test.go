package httpclient

import "testing"

func TestUrlBuilder(t *testing.T) {
	builder := &DefaultUrlBuilder{}
	builder.SetSchema("http").SetHost("192.168.1.123").SetPort("8089").SetPath("/api/user/login")
	builder.AddQueryParam("username", "admin").AddQueryParam("password", "123456")
	got := builder.Build()
	want := []string{"http://192.168.1.123:8089/api/user/login?username=admin&password=123456", "http://192.168.1.123:8089/api/user/login?password=123456&username=admin"}
	for _, url := range want {
		if url == got {
			return
		}
	}
	t.Errorf("%s", got)
}
