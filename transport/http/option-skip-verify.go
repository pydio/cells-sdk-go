package http

import (
	"crypto/tls"
	"net/http"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
)

func WithSkipVerify(skip bool) cellssdk.Option {
	return func(t *http.Transport) *http.Transport {
		if !skip {
			return t
		}
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		return t
	}
}
