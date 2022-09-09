package http

import (
	"net/http"

	cells_sdk "github.com/pydio/cells-sdk-go/v4"
	"github.com/pydio/cells-sdk-go/v4/transport"
)

// GetClient provides an option to rather use an HTTP client that ignores SSL certificate issues.
func GetClient(sdkConfig *cells_sdk.SdkConfig) *http.Client {
	return &http.Client{Transport: transport.New(
		transport.WithSkipVerify(sdkConfig.SkipVerify),
		transport.WithCustomHeaders(sdkConfig.CustomHeaders),
	)}
}
