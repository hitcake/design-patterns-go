package httpclient

import (
	"io"
	"net/http"
)

type HttpRequestBuilder interface {
	AddHeader(key, value string) HttpRequestBuilder
	SetMethod(method string) HttpRequestBuilder
	SetUrl(url string) HttpRequestBuilder
	SetBody(body io.Reader) HttpRequestBuilder
	Build() (*http.Request, error)
}

type DefaultHttpRequestBuilder struct {
	headers map[string]string
	method  string
	url     string
	body    io.Reader
}

func (builder *DefaultHttpRequestBuilder) AddHeader(key, value string) HttpRequestBuilder {
	if builder.headers == nil {
		builder.headers = make(map[string]string)
	}
	builder.headers[key] = value
	return builder
}

func (builder *DefaultHttpRequestBuilder) SetMethod(method string) HttpRequestBuilder {
	builder.method = method
	return builder
}

func (builder *DefaultHttpRequestBuilder) SetUrl(url string) HttpRequestBuilder {
	builder.url = url
	return builder
}

func (builder *DefaultHttpRequestBuilder) SetBody(body io.Reader) HttpRequestBuilder {
	builder.body = body
	return builder
}
func (builder *DefaultHttpRequestBuilder) Build() (*http.Request, error) {
	req, err := http.NewRequest(builder.method, builder.url, builder.body)
	if err != nil {
		return nil, err
	}
	for key, value := range builder.headers {
		req.Header.Set(key, value)
	}
	return req, nil
}
