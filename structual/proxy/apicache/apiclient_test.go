package apicache

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestApiClient(t *testing.T) {
	client := &http.Client{Timeout: 10 * time.Second}
	var apiClient ApiClient
	apiClient = &DefaultApiClient{client}
	apiClient = NewApiClientProxy(apiClient, 10)
	req, err := http.NewRequest("GET", "https://www.baidu.com/", nil)
	if err != nil {
		t.Error("build request error", err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			apiClient.Do(req)
		}()
	}
	wg.Wait()
	fmt.Println("wait for expire")
	time.Sleep(10 * time.Second) // 等待过期
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			apiClient.Do(req)
		}()
	}
	wg.Wait()
}
