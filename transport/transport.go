package transport

import (
	"net/http"

	cells_sdk "github.com/pydio/cells-sdk-go/v4"
)

const (
	CellsApiResourcePath = "/a"
	CellsS3SecretDefault = "gatewaysecret"

	UserAgentKey = "User-Agent"
)

type Option func(t *http.Transport) *http.Transport
type RoundTripOption func(t http.RoundTripper) http.RoundTripper

type ConfigStore interface {
	RefreshIfRequired(*cells_sdk.SdkConfig) (bool, error)
}

type TokenProvider interface {
	Retrieve() (string, error)
	Expired() bool
}

func New(options ...interface{}) http.RoundTripper {
	// First go through Transport options
	baseT := &http.Transport{}
	for _, o := range options {
		switch typed := o.(type) {
		case Option:
			baseT = typed(baseT)
		}
	}

	// Now use transport as a RoundTripper and go through RoundTripOptions
	var roundTrip http.RoundTripper
	roundTrip = baseT
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
