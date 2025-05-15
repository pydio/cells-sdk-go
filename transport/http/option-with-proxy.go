package http

import (
	"net/http"
	"net/url"

	cellssdk "github.com/pydio/cells-sdk-go/v4"
)

func WithProxy(fixedURL *url.URL) cellssdk.HttpClientOption {
	return func(t *http.Client) *http.Client {
		t.Transport = &http.Transport{Proxy: http.ProxyURL(fixedURL)}
		return t
	}
}
