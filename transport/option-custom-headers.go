package transport

import (
	"net/http"
)

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

func WithCustomHeaders(h map[string]string) RoundTripOption {
	return func(t http.RoundTripper) http.RoundTripper {
		// if h == nil || len(h) == 0 {
		if len(h) == 0 {
			return t
		}
		return &headerRoundTripper{
			rt:      t,
			Headers: h,
		}
	}
}
