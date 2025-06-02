package transport

import (
	"crypto/tls"
	"net/http"

	"github.com/pydio/cells-sdk-go/v5/apiv2"
)

func WithSkipVerify(skip bool) apiv2.TransportOption {
	return func(t *http.Transport) *http.Transport {
		if !skip {
			return t
		}
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		return t
	}
}
