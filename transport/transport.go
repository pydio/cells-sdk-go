package transport

import (
	"net/http"
	"time"

	cellssdk "github.com/pydio/cells-sdk-go/v4"
)

// New creates a new http transport with the passed round trip options and reasonable defaults.
func New(options ...any) http.RoundTripper {

	newTransport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	//	for _, o := range options {
	//		switch typed := o.(type) {
	//		case cellssdk.TransportOption:
	//			newTransport = typed(newTransport)
	//		}
	//	}

	// Cast as more generic RoundTRipper and apply corresponding RoundTripOptions
	var roundTrip http.RoundTripper
	roundTrip = newTransport
	for _, o := range options {
		switch typed := o.(type) {
		case cellssdk.RoundTripOption:
			roundTrip = typed(roundTrip)
		}
	}
	return roundTrip
}

func TokenProviderFromConfig(c *cellssdk.SdkConfig) (cellssdk.TokenProvider, error) {
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
