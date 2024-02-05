package http

import (
	"net/http"
	"time"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
)

func WithTimout(timeout time.Duration) cellssdk.HttpClientOption {
	return func(t *http.Client) *http.Client {
		t.Timeout = timeout
		return t
	}
}
