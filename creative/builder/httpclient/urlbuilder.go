package httpclient

import (
	"fmt"
	"net/url"
)

type UrlBuilder interface {
	SetSchema(schema string) UrlBuilder
	SetHost(host string) UrlBuilder
	SetPort(port string) UrlBuilder
	SetPath(path string) UrlBuilder
	AddQueryParam(key, value string) UrlBuilder
	Build() string
}

type DefaultUrlBuilder struct {
	scheme string
	host   string
	port   string
	path   string
	query  url.Values
}

func (builder *DefaultUrlBuilder) SetSchema(schema string) UrlBuilder {
	builder.scheme = schema
	return builder
}

func (builder *DefaultUrlBuilder) SetHost(host string) UrlBuilder {
	builder.host = host
	return builder
}
func (builder *DefaultUrlBuilder) SetPort(port string) UrlBuilder {
	builder.port = port
	return builder
}

func (builder *DefaultUrlBuilder) SetPath(path string) UrlBuilder {
	builder.path = path
	return builder
}
func (builder *DefaultUrlBuilder) AddQueryParam(key, value string) UrlBuilder {
	if builder.query == nil {
		builder.query = url.Values{}
	}
	builder.query.Add(key, value)
	return builder
}
func (builder *DefaultUrlBuilder) Build() string {
	if builder.scheme == "" {
		builder.scheme = "http"
	}
	if builder.host == "" {
		builder.host = "localhost"
	}
	if builder.port == "" {
		builder.port = "80"
	}
	if builder.path == "" {
		builder.path = "/"
	}
	return fmt.Sprintf("%s://%s:%s%s?%s", builder.scheme, builder.host, builder.port, builder.path, builder.query.Encode())
}
