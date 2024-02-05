package http

import (
	"context"
	"net/http"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
)

func WithBearer(provider cellssdk.TokenProvider) cellssdk.RoundTripOption {
	return func(t http.RoundTripper) http.RoundTripper {
		return &bearerRoundTripper{rt: t, tp: provider}
	}
}

type bearerRoundTripper struct {
	rt http.RoundTripper
	tp cellssdk.TokenProvider
}

func (r *bearerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if token, er := r.tp.Retrieve(context.TODO()); er == nil {
		req.Header.Set("Authorization", "Bearer "+token)
		return r.rt.RoundTrip(req)
	} else {
		return nil, er
	}
}
