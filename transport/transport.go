package transport

import (
	"net/http"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
)

const (
	CellsApiResourcePath = "/a"
	CellsS3SecretDefault = "gatewaysecret"

	UserAgentKey = "User-Agent"
)

type Option func(t *http.Transport) *http.Transport

type RoundTripOption func(t http.RoundTripper) http.RoundTripper

// New creates a new default http transport with the passed transport and roundtrip options.
func New(options ...interface{}) http.RoundTripper {

	// Creates a new default http transport and applies relevant transport options
	newTransport := &http.Transport{}
	for _, o := range options {
		switch typed := o.(type) {
		case Option:
			newTransport = typed(newTransport)
		}
	}

	// Cast as more generic RoundTRipper and apply corresponding RoundTripOptions
	var roundTrip http.RoundTripper
	roundTrip = newTransport
	for _, o := range options {
		switch typed := o.(type) {
		case RoundTripOption:
			roundTrip = typed(roundTrip)
		}
	}
	return roundTrip
}

func TokenProviderFromConfig(c *cells_sdk.SdkConfig) (TokenProvider, error) {
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
