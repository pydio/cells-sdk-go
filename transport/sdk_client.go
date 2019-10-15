package transport

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pydio/cells-sdk-go/transport/oidc"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	cells_sdk "github.com/pydio/cells-sdk-go"
	http2 "github.com/pydio/cells-sdk-go/transport/http"
)

var (
	apiResourcePath = "/a"
)

func GetRestClientTransport(sdkConfig *cells_sdk.SdkConfig, anonymous bool) (context.Context, runtime.ClientTransport, error) {
	u, e := url.Parse(sdkConfig.Url)
	if e != nil {
		return nil, nil, e
	}

	ctx := context.Background()
	transport := httptransport.New(u.Host, apiResourcePath, []string{u.Scheme})
	if sdkConfig.SkipVerify {
		transport.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired and self-signed SSL certificates
		}
	}
	if sdkConfig.CustomHeaders != nil && len(sdkConfig.CustomHeaders) > 0 {
		transport.Transport = &customHeaderRoundTripper{
			rt:      transport.Transport,
			Headers: sdkConfig.CustomHeaders,
		}
	}
	transport.Context = ctx

	if anonymous {
		return ctx, transport, nil
	}

	jwt, err := oidc.RetrieveToken(sdkConfig)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"cannot retrieve token with config:\n%s - %s - %s - %s - %s - %v\nerror cause: %s",
			sdkConfig.Url, sdkConfig.ClientKey, sdkConfig.ClientSecret,
			sdkConfig.User, sdkConfig.Password, sdkConfig.SkipVerify, err.Error())
	}
	bearerTokenAuth := httptransport.BearerToken(jwt)
	transport.DefaultAuthentication = bearerTokenAuth

	return ctx, transport, nil
}

type customHeaderRoundTripper struct {
	rt      http.RoundTripper
	Headers map[string]string
}

func (c customHeaderRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	return c.rt.RoundTrip(req)
}
