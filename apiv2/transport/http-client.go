package transport

import (
	cellsSdk "github.com/pydio/cells-sdk-go/v5/apiv2"
	"net/http"
)

// GetClient provides an option to rather use an HTTP client that ignores SSL certificate issues.
func GetClient(sdkConfig *cellsSdk.SdkConfig) *http.Client {
	return &http.Client{Transport: New(
		WithSkipVerify(sdkConfig.SkipVerify),
		WithCustomHeaders(sdkConfig.CustomHeaders),
	)}
}
