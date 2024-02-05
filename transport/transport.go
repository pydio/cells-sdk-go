package transport

import (
	"net/http"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
)

const (
	CellsApiResourcePath = "/a"
	UserAgentKey         = "User-Agent"
)

// New creates a new default http transport with the passed transport and round-trip options.
func New(options ...interface{}) http.RoundTripper {

	// Creates a new default http transport and applies relevant transport options
	newTransport := &http.Transport{}
	for _, o := range options {
		switch typed := o.(type) {
		case cellssdk.Option:
			newTransport = typed(newTransport)
		}
	}

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
