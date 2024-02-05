package transport

import (
	"net/http"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
	http2 "github.com/pydio/cells-sdk-go/v5/transport/http"
)

// NewHttpClient creates a custom HTTP client that has correct TLS skip verify flag and
// defined custom headers.
// Caller can provide additional options, typically to set a custom timeout.
func NewHttpClient(sdkConfig *cellssdk.SdkConfig, options ...interface{}) *http.Client {
	client := &http.Client{
		Transport: New(
			http2.WithSkipVerify(sdkConfig.SkipVerify),
			http2.WithCustomHeaders(sdkConfig.CustomHeaders),
		),
	}
	// browse generic options to find the relevant ones, e.G to define a custom timeout.
	for _, o := range options {
		switch typed := o.(type) {
		case cellssdk.HttpClientOption:
			client = typed(client)
		}
	}
	return client
}
