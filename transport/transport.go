package transport

import (
	"net/http"

	cells_sdk "github.com/pydio/cells-sdk-go/v4"
)

const (
	CellsApiResourcePath = "/a"
	CellsS3SecretDefault = "gatewaysecret"
)

type Option func(t *http.Transport) *http.Transport
type RoundTripOption func(t http.RoundTripper) http.RoundTripper

type TokenProvider interface {
	Retrieve() (string, error)
	Expired() bool
}

func New(options ...interface{}) http.RoundTripper {
	// First go through Transport options
	baseT := &http.Transport{}
	for _, o := range options {
		switch o.(type) {
		case Option:
			to := o.(Option)
			baseT = to(baseT)
		}
	}

	// Now use transport as a RoundTripper and go through RoundTripOptions
	var roundTrip http.RoundTripper
	roundTrip = baseT
	for _, o := range options {
		switch o.(type) {
		case RoundTripOption:
			to := o.(RoundTripOption)
			roundTrip = to(roundTrip)
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
