package http

import (
	cellssdk "github.com/pydio/cells-sdk-go/v4/apiv1"
	"net/http"
	"time"
)

func WithTimout(timeout time.Duration) cellssdk.HttpClientOption {
	return func(t *http.Client) *http.Client {
		t.Timeout = timeout
		return t
	}
}
