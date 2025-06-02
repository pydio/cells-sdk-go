package http

import (
	"net/http"
	"time"

	"github.com/pydio/cells-sdk-go/v5/apiv1"
)

func WithTimout(timeout time.Duration) apiv1.HttpClientOption {
	return func(t *http.Client) *http.Client {
		t.Timeout = timeout
		return t
	}
}
