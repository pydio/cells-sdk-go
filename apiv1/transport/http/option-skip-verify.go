package http

import (
	"crypto/tls"
	cellssdk "github.com/pydio/cells-sdk-go/v4/apiv1"
	"net/http"
)

func WithSkipVerify(skip bool) cellssdk.TransportOption {
	return func(t *http.Transport) *http.Transport {
		if !skip {
			return t
		}
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		return t
	}
}
