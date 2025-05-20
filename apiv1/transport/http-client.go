package transport

import (
	"net/http"

	"github.com/pydio/cells-sdk-go/v5/apiv1"
	http3 "github.com/pydio/cells-sdk-go/v5/apiv1/transport/http"
)

// NewHttpClient creates a custom HTTP client that has correct TLS skip verify flag and
// defined custom headers.
// Caller can provide additional options, typically to set a custom timeout.
func NewHttpClient(sdkConfig *apiv1.SdkConfig, options ...any) *http.Client {
	client := &http.Client{
		Transport: New(
			http3.WithSkipVerify(sdkConfig.SkipVerify),
			http3.WithCustomHeaders(sdkConfig.CustomHeaders),
		),
	}
	// Apply relevant options, e.G to define a custom timeout.
	for _, o := range options {
		switch typed := o.(type) {
		case apiv1.HttpClientOption:
			client = typed(client)
		}
	}
	return client
}
