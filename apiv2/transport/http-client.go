package transport

import (
	"net/http"

	"github.com/pydio/cells-sdk-go/v5/apiv2"
)

// GetClient provides an option to rather use an HTTP client that ignores SSL certificate issues.
func GetClient(sdkConfig *apiv2.SdkConfig) *http.Client {
	return &http.Client{Transport: New(
		WithSkipVerify(sdkConfig.SkipVerify),
		WithCustomHeaders(sdkConfig.CustomHeaders),
	)}
}
