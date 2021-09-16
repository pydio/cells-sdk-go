package http

import (
	"net/http"

	cells_sdk "github.com/pydio/cells-sdk-go/v3"
	"github.com/pydio/cells-sdk-go/v3/transport"
)

// GetClient provides an option to rather use an HTTP client that ignores SSL certificate issues.
func GetClient(sdkConfig *cells_sdk.SdkConfig) *http.Client {
	return &http.Client{Transport: transport.New(
		transport.WithSkipVerify(sdkConfig.SkipVerify),
		transport.WithCustomHeaders(sdkConfig.CustomHeaders),
	)}
}
