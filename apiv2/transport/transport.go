package transport

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"

	openapiRuntime "github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/apiv2"
)

const (
	CellsApiPrefix       = "/v2"
	CellsS3SecretDefault = "gatewaysecret"

	KeyUserAgent = "User-Agent"
)

func BasicAuthWriter(ctx context.Context, currConfig *apiv2.SdkConfig) openapiRuntime.ClientAuthInfoWriter {

	return openapiRuntime.ClientAuthInfoWriterFunc(func(r openapiRuntime.ClientRequest, _ strfmt.Registry) error {
		return fmt.Errorf("unsupported authentication mode for Cells API v2")
	})

	//currTransport, err := GetRuntimeTransport(ctx, currConfig)
	//if err != nil {
	//	log.Fatal("cannot get runtime transport", err)
	//}

	// TODO we do not have this in the API v2 yet
	// 	frontendService := sdkClient.New(currTransport, strfmt.Default).FrontendService
	//
	//frontend_service.NewFrontSessionParams().WithBody(&models.RestFrontSessionRequest{
	//	AuthInfo: map[string]string{
	//		"login":    f.user,
	//		"password": f.password,
	//		"type":     "credentials",
	//	},
	//}).WithContext(ctx)
}

func UserAgent() string {
	osVersion := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	goVersion := fmt.Sprintf("%s", runtime.Version())
	appVersion := fmt.Sprintf("pydio/cells-sdk-go@v%s", "5-dev")
	return fmt.Sprintf("%s %s %s", osVersion, goVersion, appVersion)
}

type Option func(t *http.Transport) *http.Transport
type RoundTripOption func(t http.RoundTripper) http.RoundTripper

type TokenProvider interface {
	Retrieve() (string, error)
	Expired() bool
}

// New creates a new http transport with the passed round-trip options and reasonable defaults.
func New(options ...interface{}) http.RoundTripper {
	// First go through Transport options
	newTransport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// Cast as more generic RoundTRipper and apply corresponding RoundTripOptions
	var roundTrip http.RoundTripper
	roundTrip = newTransport
	for _, o := range options {
		switch o.(type) {
		case RoundTripOption:
			to := o.(RoundTripOption)
			roundTrip = to(roundTrip)
		}
	}

	return roundTrip
}

func TokenProviderFromConfig(c *apiv2.SdkConfig) (TokenProvider, error) {
	// TODO we only support PAT for the time being
	return c, nil // SdkConfig implements TokenProvider interface
	//if c.IdToken != "" {
	//	return c, nil // SdkConfig implements TokenProvider interface
	//} else {
	//	tp, e := NewFrontSessionTokenProvider(c)
	//	if e != nil {
	//		return nil, e
	//	}
	//	if c.UseTokenCache {
	//		tp = WithProviderCache(tp, c)
	//	}
	//	return tp, nil
	//}
}
