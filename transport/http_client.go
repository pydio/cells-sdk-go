package transport

import (
	"net/http"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
)

// GetClient provides an option to rather use an HTTP client that ignores SSL certificate issues.
func GetClient(sdkConfig *cells_sdk.SdkConfig) *http.Client {
	return &http.Client{Transport: New(
		WithSkipVerify(sdkConfig.SkipVerify),
		WithCustomHeaders(sdkConfig.CustomHeaders),
	)}
}
