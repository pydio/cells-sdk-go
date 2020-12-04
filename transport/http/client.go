package http

import (
	"net/http"

	"github.com/pydio/cells-sdk-go/transport"

	"github.com/pydio/cells-sdk-go"
)

// GetClient provides an option to rather use an http client that ignore SSL certificate issues.
func GetClient(sdkConfig *cells_sdk.SdkConfig) *http.Client {
	return &http.Client{Transport: transport.New(
		transport.WithSkipVerify(sdkConfig.SkipVerify),
		transport.WithCustomHeaders(sdkConfig.CustomHeaders),
	)}
}
