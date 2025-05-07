package transport

import (
	"net/http"
)

func WithBearer(provider TokenProvider) RoundTripOption {
	return func(t http.RoundTripper) http.RoundTripper {
		return &bearerRoundTripper{rt: t, tp: provider}
	}
}

type bearerRoundTripper struct {
	rt http.RoundTripper
	tp TokenProvider
}

func (r *bearerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if token, er := r.tp.Retrieve(); er == nil {
		req.Header.Set("Authorization", "Bearer "+token)
		return r.rt.RoundTrip(req)
	} else {
		return nil, er
	}
}
