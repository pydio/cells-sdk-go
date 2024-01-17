package http

import (
	"net/http"
	"time"

	cells_sdk "github.com/pydio/cells-sdk-go/v4"
	"github.com/pydio/cells-sdk-go/v4/transport"
)

// GetClient provides an option to rather use an HTTP client that ignores SSL certificate issues.
func GetClient(sdkConfig *cells_sdk.SdkConfig) *http.Client {
	return &http.Client{
		Transport: transport.New(
			transport.WithSkipVerify(sdkConfig.SkipVerify),
			transport.WithCustomHeaders(sdkConfig.CustomHeaders),
		),
	}
}

// GetClient provides an option to rather use an HTTP client that ignores SSL certificate issues.
func GetClientWithTimeout(sdkConfig *cells_sdk.SdkConfig, timeout time.Duration) *http.Client {
	return &http.Client{
		Transport: transport.New(
			transport.WithSkipVerify(sdkConfig.SkipVerify),
			transport.WithCustomHeaders(sdkConfig.CustomHeaders),
		),
		// TODO verify and if correct define as a parameter of the config
		Timeout: timeout,
	}
}
