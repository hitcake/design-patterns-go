package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserToken struct {
	Token string `json:"token"`
}
type UserInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin UserLogin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userLogin); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		if userLogin.Username == "admin" && userLogin.Password == "123456" {
			w.WriteHeader(http.StatusBadRequest)
			userToken := UserToken{Token: "token"}
			response := Response{Code: 0, Msg: "success", Data: userToken}
			_ = json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusOK)
			response := Response{Code: -1, Msg: "error"}
			_ = json.NewEncoder(w).Encode(response)
		}

	}
}

func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		userInfo := UserInfo{Username: "admin", Name: "admin", Role: "administrator"}
		response := Response{Code: 0, Msg: "success", Data: userInfo}
		json.NewEncoder(w).Encode(response)
	}
}

func startServer() {
	const token = "123456abcdef"
	http.HandleFunc("/api/user/profile", userProfileHandler)
	http.HandleFunc("/api/user/login", userLoginHandler)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			return
		}
	}()
	time.Sleep(1 * time.Second)
}

func TestHttpRequestBuilder(t *testing.T) {
	startServer()
	token := userLogin(t)
	userProfile(t, token)
}

func userLogin(t *testing.T) string {
	urlBuilder := &DefaultUrlBuilder{}
	urlBuilder.SetSchema("http").SetHost("127.0.0.1").SetPort("8080").SetPath("/api/user/login")
	url := urlBuilder.Build()

	requestBuilder := &DefaultHttpRequestBuilder{}
	loginMap := map[string]string{}
	loginMap["username"] = "admin"
	loginMap["password"] = "123456"
	b, err := json.Marshal(loginMap)
	if err != nil {
		t.Errorf("Error marshalling body: %v", err)
	}
	requestBuilder.SetMethod(http.MethodPost).SetUrl(url)
	requestBuilder.SetBody(bytes.NewReader(b))
	requestBuilder.AddHeader("Content-Type", "application/json")
	req, err := requestBuilder.Build()
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	response := Response{}
	json.NewDecoder(resp.Body).Decode(&response)
	if response.Code != 0 {
		t.Errorf("Unexpected status code: %d", response.Code)
	}
	return response.Data.(map[string]interface{})["token"].(string)
}

func userProfile(t *testing.T, token string) {
	urlBuilder := &DefaultUrlBuilder{}
	urlBuilder.SetSchema("http").SetHost("127.0.0.1").SetPort("8080").SetPath("/api/user/profile")
	url := urlBuilder.Build()
	requestBuilder := &DefaultHttpRequestBuilder{}
	requestBuilder.AddHeader("Authorization", "Bearer "+token)
	requestBuilder.SetMethod(http.MethodGet).SetUrl(url)
	request, err := requestBuilder.Build()
	if err != nil {
		t.Error(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	response := Response{}
	json.NewDecoder(resp.Body).Decode(&response)
	if response.Code != 0 {
		t.Errorf("Unexpected status code: %d", response.Code)
	}
	role := response.Data.(map[string]interface{})["role"].(string)
	if role != "administrator" {
		t.Errorf("Unexpected role %s", role)
	}
}
