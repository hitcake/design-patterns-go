package apicache

import (
	"fmt"
	"io"
	"net/http"
)

type ApiClient interface {
	Do(*http.Request) ([]byte, error)
}
type DefaultApiClient struct {
	httpClient *http.Client
}

func (api *DefaultApiClient) Do(req *http.Request) ([]byte, error) {
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("request for %s\n", req.URL.String())
	if err != nil {
		return nil, err
	}
	return body, nil
}
