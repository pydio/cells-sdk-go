package transport

import (
	"crypto/tls"
	"net/http"
)

func WithSkipVerify(skip bool) Option {
	return func(t *http.Transport) *http.Transport {
		if !skip {
			return t
		}
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		return t
	}
}
