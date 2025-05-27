package transport

import (
	"net/http"

	"github.com/pydio/cells-sdk-go/v5/apiv2"
)

func WithCustomHeaders(h map[string]string) apiv2.RoundTripOption {
	return func(t http.RoundTripper) http.RoundTripper {
		if len(h) == 0 {
			return t
		}
		return &headerRoundTripper{
			rt:      t,
			Headers: h,
		}
	}
}

type headerRoundTripper struct {
	rt      http.RoundTripper
	Headers map[string]string
}

func (c headerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	return c.rt.RoundTrip(req)
}
