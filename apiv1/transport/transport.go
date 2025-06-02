package transport

import (
	"net/http"
	"time"

	"github.com/pydio/cells-sdk-go/v5/apiv1"
)

// New creates a new default http transport with the passed transport and round-trip options.
func New(options ...any) http.RoundTripper {

	// Creates a new default http transport and applies relevant transport options
	newTransport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	for _, o := range options {
		switch typed := o.(type) {
		case apiv1.TransportOption:
			newTransport = typed(newTransport)
		}
	}

	// Cast as more generic RoundTripper and apply corresponding RoundTripOptions
	var roundTrip http.RoundTripper
	roundTrip = newTransport
	for _, o := range options {
		switch typed := o.(type) {
		case apiv1.RoundTripOption:
			roundTrip = typed(roundTrip)
		}
	}
	return roundTrip
}

func TokenProviderFromConfig(c *apiv1.SdkConfig) (apiv1.TokenProvider, error) {
	if c.IdToken != "" {
		return c, nil // SdkConfig implements TokenProvider interface
	} else {
		tp, e := NewFrontSessionTokenProvider(c)
		if e != nil {
			return nil, e
		}
		if c.UseTokenCache {
			tp = WithProviderCache(tp, c)
		}
		return tp, nil
	}
}
