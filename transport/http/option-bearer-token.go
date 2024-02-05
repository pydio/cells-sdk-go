package http

import (
	"net/http"

	"github.com/pydio/cells-sdk-go/v5"
)

func WithBearer(provider cells_sdk.TokenProvider) cells_sdk.RoundTripOption {
	return func(t http.RoundTripper) http.RoundTripper {
		return &bearerRoundTripper{rt: t, tp: provider}
	}
}

type bearerRoundTripper struct {
	rt http.RoundTripper
	tp cells_sdk.TokenProvider
}

func (r *bearerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if token, er := r.tp.Retrieve(); er == nil {
		req.Header.Set("Authorization", "Bearer "+token)
		return r.rt.RoundTrip(req)
	} else {
		return nil, er
	}
}