package http

import (
	"net/http"
	"time"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
)

func WithTimout(timeout time.Duration) cells_sdk.HttpClientOption {
	return func(t *http.Client) *http.Client {
		t.Timeout = timeout
		return t
	}
}
